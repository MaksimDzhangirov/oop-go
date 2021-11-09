# Абстрактные классы в Go

[Оригинал](https://golangbyexample.com/go-abstract-class/)

Интерфейс в Go не содержит полей, а также не позволяет определять методы внутри
него. Любой тип должен реализовывать все методы интерфейса, чтобы иметь тип
этого интерфейса. Существуют ситуации, когда полезно иметь реализацию метода по
умолчанию, а также поля по умолчанию в Go. Прежде чем понять как это можно 
сделать давайте сначала разберёмся с требованиями для абстрактного класса:

1. Абстрактный класс должен иметь поля по умолчанию
2. Абстрактный класс должен иметь метод по умолчанию
3. Не должно быть возможности непосредственно создать экземпляр абстрактного 
   класса
   
Мы будем использовать комбинацию интерфейса (абстрактный интерфейс) и структуры
(абстрактный конкретный тип). С их помощью можно будет реализовать функционал 
абстрактного класса. Смотри программу ниже:

```go
package main

import "fmt"

// Абстрактный интерфейс
type iAlpha interface {
    work()
    common()
}

// Абстрактный конкретный тип
type alpha struct {
    name string
}

func (a *alpha) common() {
    fmt.Println("common called")
}

// Реализуем тип
type beta struct {
    alpha
}

func (b *beta) work() {
    fmt.Println("work called")
    fmt.Printf("name is %s\n", b.name)
    b.common()
}

func main() {
    a := alpha{
        name: "test",
    }
    b := &beta{
        alpha: a,
    }
    b.work()
}
```

Результат в терминале:

```shell
go run abstractClass/example1/program1.go 
work called
name is test
common called
```

В вышеприведенной программе:

* мы создали абстрактный интерфейс `iAlpha`, абстрактную конкретную структуру 
  `alpha` и структуру, в которой они реализованы, `beta`.
* Структура `alpha` встроена в `beta`
* Структура `beta` имеет доступ к полю по умолчанию `name`
* Структура `beta` имеет доступ к методу по умолчанию `common`
* Невозможно использовать непосредственно структуру `alpha` не получится, поскольку
  она поддерживает только один из методов интерфейса `iAlpha`.
  
Таким образом, все требования удовлетворены, но существует одно ограничение 
вышеупомянутого способа. Невозможно вызвать метод `work` из `common` в `alpha`.
По сути нет способа вызвать неопределенный метод абстрактного интерфейса из 
методов по умолчанию абстрактного конкретного типа. Однако это можно обойти 
следующим образом:

```go
package main

import "fmt"

// Абстрактный интерфейс
type iAlpha interface {
    work()
    common()
}

// Абстрактный конкретный тип
type alpha struct {
    name string
    work func()
}

func (a *alpha) common() {
    fmt.Println("common called")
    a.work()
}

// Реализуем тип
type beta struct {
    alpha
}

func (b *beta) work() {
    fmt.Println("work called")
    fmt.Printf("name is %s\n", b.name)
}

func main() {
    a := alpha{
        name: "test",
    }
    b := &beta{
        alpha: a,
    }
    b.alpha.work = b.work
    b.common()
}
```

Результат в терминале:

```shell
go run abstractClass/example2/program2.go
common called
work called
name is test
```

В вышеприведенной программе:
* Мы создали новое поле `work` типа `func` в `alpha`
* Мы присвоили методу `work` из `alpha` метод `work` из `beta`

Единственная проблема в приведенной выше программе заключается в том, что можно
непосредственно создать экземпляр структуры `alpha` и задав определение для
метода `work` она будет удовлетворять типу `iAlpha`. Это нарушает пункт 3 требований
к абстрактному классу. Попробуем исправить эту проблему. Приведенная ниже программа
решает обе проблемы:

```go
package main

import "fmt"

// Абстрактный интерфейс
type iAlpha interface {
    work()
    common(iAlpha)
}

// Абстрактный конкретный тип
type alpha struct {
    name string
}

func (a *alpha) common(i iAlpha) {
    fmt.Println("common called")
    i.work()
}

// Реализуем тип
type beta struct {
    alpha
}

func (b *beta) work() {
    fmt.Println("work called")
    fmt.Printf("name is %s\n", b.name)
}

func main() {
    a := alpha{
        name: "test",
    }
    b := &beta{
        alpha: a,
    }
    b.common(b)
}
```

Результат в терминале:

```shell
go run abstractClass/example3/program3.go 
common called
work called
Name is test
```

В ней:

* Все методы по умолчанию принимают в качестве аргумента интерфейс `iAlpha`. Все
  неопределенные в структуре `alpha` методы будут вызываться, используя этот
  аргумент в методах по умолчанию.
  
## Заключение

Приведенная выше программа удовлетворяет всем трём требованиям абстрактного 
класса. Это один из способов имитации абстрактного класса в Go.