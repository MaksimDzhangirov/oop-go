# ООП: наследование в Golang

[Оригинал](https://golangbyexample.com/oop-inheritance-golang-complete/)

Мы попытаемся пояснить наследование в Go, сравнивая его с наследованием в Java.
Первое о чём хотелось бы упомянуть здесь - в Golang нет таких ключевых слов как
`extends` and `implements` в Java. Go предоставляет ограниченный функционал,
присущий ключевым словам `extends` and `implements` по-разному, причём каждый 
имеет свои ограничения. Прежде чем мы начнём рассматривать наследование в Go
стоит упомянуть, что:

* в Go предпочтительнее использовать композицию вместо наследования. Это позволяет
  встраивать одну структуру в другую.
* Go не поддерживает наследование типов.

Начнем с простейшего примера наследования в Go. Затем перечислим ограничения или
отсутствующий функционал. Затем мы попытаемся преодолеть ограничения или добавить
недостающие функции, пока не напишем программу, которая будет обладать всеми 
свойствами наследования доступными в Go. Итак, начнём.

Самый простой случай использования наследования — дочерний тип должен иметь 
доступ к полям и методам родительского типа. Это делается в Go посредством 
встраивания. Базовая структура встраивается в дочернюю, после чего 
базовые поля и методы могут быть напрямую доступы дочерней структуре. Смотри
код ниже: дочерняя структура может напрямую обращаться к полю `color`, а также
напрямую вызывать метод `say()`.

```go
package main

import "fmt"

type base struct {
    color string
}

func (b *base) say() {
    fmt.Println("Hi from say function")
}

type child struct {
    base  // встраиваем
    style string
}

func main() {
    base := base{color: "Red"}
    child := &child{
        base:  base,
        style: "somestyle",
    }
    child.say()
    fmt.Println("The color is " + child.color)
}
```

Результат в терминале:

```shell
go run inheritance/example1/program1.go
Hi from say function
The color is Red
```

Одним из ограничений вышеприведенной программы является то, что Вы не можете 
передать дочерний тип в функцию, которая ожидает базовый тип, поскольку Go не
допускает наследования типов. Например, приведённый ниже код не компилируется 
и выдаёт ошибку - "**cannot use child (type *child) as type base in argument to
check**".

```go
package main

import "fmt"

type base struct {
    color string
}

func (b *base) say() {
    fmt.Println("Hi from say function")
}

type child struct {
    base  // встраиваем
    style string
}

func check(b base) {
    b.say()
}

func main() {
    base := base{color: "Red"}
    child := &child{
        base:  base,
        style: "somestyle",
    }
    child.say()
    fmt.Println("The color is " + child.color)
    check(child)
}
```

Результат в терминале:

```shell
go run inheritance/example2/program2.go
inheritance/example2/program2.go:30:7: cannot use child (type *child) as type base in argument to check
```

Вышеприведенная ошибка говорит о том, что дочерний тип не может быть задан в Go
просто с помощью встраивания. Попробуем исправить эту ошибку. Для этого нам
необходимо использовать интерфейсы Go. Смотри показанную ниже версию программы,
в которой исправлена эта ошибка.

```go
package main

import "fmt"

type iBase interface {
    say()
}

type base struct {
    color string
}

func (b *base) say() {
    fmt.Println("Hi from say function")
}

type child struct {
    base  // встраиваем
    style string
}

func check(b iBase) {
    b.say()
}

func main() {
    base := base{color: "Red"}
    child := &child{
        base:  base,
        style: "somestyle",
    }
    child.say()
    fmt.Println("The color is " + child.color)
    check(child)
}
```

Результат в терминале:

```shell
go run inheritance/example3/program3.go
Hi from say function
The color is Red
Hi from say function
```

В приведенной выше программе мы: (a) создали интерфейс `iBase`, который 
содержит метод `say`; (b) изменили метод «check», чтобы он принимал в качестве 
аргумента тип `iBase`.

Поскольку базовая структура реализует метод `say`, а в дочернюю структуру в 
свою очередь встраивается базовая, то дочерняя структура косвенно реализует 
метод `say` и становится типа `iBase`. Из-за этого теперь мы можем передать её
в функцию `check`. Отлично, мы устранили одно ограничение, скомбинировав 
структуру и интерфейс.

Но существует ещё одно ограничение. Допустим у базовой и дочерней структуры 
есть ещё один метод `clear`, а метод `say` вызывает метод `clear`. Затем
когда метод `say` вызывается в дочерней структуре, он запускает метод `clear` 
базового, а не дочернего объекта. Смотри пример ниже.

```go
package main

import "fmt"

type iBase interface {
    say()
}

type base struct {
    color string
}

func (b *base) say() {
    b.clear()
}

func (b *base) clear() {
    fmt.Println("Clear from base's function")
}

type child struct {
    base  // встраиваем
    style string
}

func (c *child) clear() {
    fmt.Println("Clear from child's function")
}

func check(b iBase) {
    b.say()
}

func main() {
    base := base{color: "Red"}
    child := &child{
        base:  base,
        style: "somestyle",
    }
    child.say()
}
```

Результат в терминале:

```shell
go run inheritance/example4/program4.go
Clear from base's function
```

Как видите, вместо дочернего метода вызывается базовый метод `clear`. Это поведение
отличается от того, что происходит в Java, где был бы вызван метод `clear`
объекта `child`.

Один из способов решить указанную выше проблему — сделать `clear` полем в базовой
структуру типа `function`. Это возможно в Go, поскольку в нём функции являются
объектами первого класса. Ниже показано решение.

```go
package main

import "fmt"

type iBase interface {
    say()
}
type base struct {
    color string
    clear func()
}
func (b *base) say() {
    b.clear()
}
type child struct {
    base  //embedding
    style string
}
func check(b iBase) {
    b.say()
}
func main() {
    base := base{color: "Red",
        clear: func() {
            fmt.Println("Clear from child's function")
        }}
    child := &child{
        base:  base,
        style: "somestyle",
    }
    child.say()
}
```

Результат в терминале:

```shell
go run inheritance/example5/program5.go
Clear from child's function
```

Давайте попробуем добавить ещё один функционал в вышеприведенную программу, а
именно:

Множественное наследование — дочерняя структура должна иметь доступ к полям и
методам двух базовых структур, а также должна быть возможна подтипизация. Вот 
пример такого кода:

```go
package main

import "fmt"

type iBase1 interface {
	say()
}

type iBase2 interface {
	walk()
}

type base1 struct {
	color string
}

func (b *base1) say() {
	fmt.Println("Hi from say function")
}

type base2 struct {
}

func (b *base1) walk() {
	fmt.Println("Hi from walk function")
}

type child struct {
	base1 // встраиваем
	base2 // встраиваем
	style string
}

func check1(b iBase1) {
	b.say()
}

func check2(b iBase2) {
	b.walk()
}

func main() {
	base1 := base1{
		color: "Red",
	}
	base2 := base2{}
	child := &child{
		base1: base1,
		base2: base2,
		style: "somestyle",
	}
	child.say()
	child.walk()
	check1(child)
	check2(child)
}
```

Результат в терминале:

```shell
go run inheritance/example6/program6.go
Hi from say function
Hi from walk function
Hi from say function
Hi from walk function
```

В вышеприведенной программе в дочернюю структуру встраивается как `base1`, так и
`base2`. Его также можно передать как экземпляр интерфейса `iBase1` и `iBase2` 
функциям `check1` и `check2` соответственно. Так мы добиваемся множественного 
наследования.

Теперь главный вопрос как реализовать "Иерархию типов" в Go. Как уже было сказано
выше, наследование типов не допускается и, следовательно, в Go нет иерархии типов.
В Go намеренно это не реализовано, поэтому любое изменение в поведении 
интерфейса распространяется только на те структуры, которые ему удовлетворяют.

Тем не менее, мы можем реализовать иерархию типов, используя интерфейсы и 
структуру, как показано ниже.

```go
package main
import "fmt"
type iAnimal interface {
    breathe()
}
type animal struct {
}
func (a *animal) breathe() {
    fmt.Println("Animal breate")
}
type iAquatic interface {
    iAnimal
    swim()
}
type aquatic struct {
    animal
}
func (a *aquatic) swim() {
    fmt.Println("Aquatic swim")
}
type iNonAquatic interface {
    iAnimal
    walk()
}
type nonAquatic struct {
    animal
}
func (a *nonAquatic) walk() {
    fmt.Println("Non-Aquatic walk")
}
type shark struct {
    aquatic
}
type lion struct {
    nonAquatic
}
func main() {
    shark := &shark{}
    checkAquatic(shark)
    checkAnimal(shark)
    lion := &lion{}
    checkNonAquatic(lion)
    checkAnimal(lion)
}
func checkAquatic(a iAquatic) {}
func checkNonAquatic(a iNonAquatic) {}
func checkAnimal(a iAnimal) {}
```

Посмотрите как в вышеприведенной программе мы смогли создать иерархию. Это 
характерный для Go способ создания иерархии типов с использованием встраивания как
на уровне структуры, так и на уровне интерфейса. Здесь следует отметить, что
если нужно различать типы, то есть `shark` не должна быть одновременно `iAquatic`
и `iNonAquatic`, тогда должен быть хотя бы один метод в наборе методов `iAquatic`
и `iNonAquatic`, которого нет в другом. В нашем примере такими методами 
являются `swim` и `walk`.

Результат в терминале:

```shell
go run inheritance/example7/program7.go
iAnimal
--iAquatic
----shark
--iNonAquatic
----lion
```

## Заключение

Go не поддерживает наследование типов, но того же можно добиться с помощью 
встраивания, но нужно быть внимательным при создании такой иерархии типов. 
Кроме того, Go не поддерживает переопределение метода.
