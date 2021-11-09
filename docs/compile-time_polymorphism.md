# Полиморфизм во время компиляции в Go

[Оригинал](https://golangbyexample.com/compile-time-polymorphism-go)

При полиморфизме во время компиляции какую функцию вызывать решает компилятор.
Примерами полиморфизма во время компиляции могут быть:
* перегрузка метода/функции: существует более одного метода/функции с одним и 
  тем же именем, но с разными сигнатурами или, возможно, с разными типами 
  возвращаемых значений;
* перегрузка операторов: один и тот же оператор используется для работы с 
  разными типами данных.

Go не поддерживает перегрузку метода. Это можно показать с помощью следующей 
программы:

```go
package main

import "fmt"

type maths struct {
}

func (m *maths) add(a, b int) int {
	return a + b
}

func (m *maths) add(a, b, c int) int {
	return a + b + c
}

func main() {
	m := &maths{}
	fmt.Println(m.add(1, 2))
}
```

```shell
go run polymorphism/example1/program1.go
polymorphism/example1/program1.go:12:6: method redeclared: maths.add
        method(*maths) func(int, int) int
        method(*maths) func(int, int, int) int
polymorphism/example1/program1.go:12:17: (*maths).add redeclared in this block
        previous declaration at polymorphism/example1/program1.go:8:6
```

Go также не поддерживает перегрузку оператора. Причина указана в часто 
задаваемых вопросах - [https://golang.org/doc/faq#overloading](https://golang.org/doc/faq#overloading)

> Вызов методов упрощается, если не требуется выполнять сопоставление типов. Опыт
> работы с другими языками показал нам, что иногда полезно иметь различные 
> методы с одинаковыми именами, но с разными сигнатурами, но на практике это 
> сбивает с толку и ненадежно. Сопоставление только по имени и требование 
> согласованности типов было основным упрощающим решением в системе типов Go.
> 
> Что касается перегрузки операторов, это кажется скорее удобством, чем 
> абсолютным требованием. Опять же, без него всё проще.

Возникает вопрос: существует ли какая-нибудь альтернатива перегрузке методов в 
Go? Здесь на помощь приходят **вариативные функции**. Смотри программу ниже:

```go
package main

import "fmt"

type maths struct {
}

func (m *maths) add(numbers ...int) int {
    result := 0
    for _, num := range numbers {
        result += num
    }
    return result
}

func main() {
    m := &maths{}
  
    fmt.Printf("Result: %d\n", m.add(2, 3))
    fmt.Printf("Result: %d\n", m.add(2, 3, 4))
}
```

Результат в терминале:

```shell
go run polymorphism/example2/program2.go
Result: 5
Result: 9
```

# Заключение

Go не поддерживает напрямую перегрузку методов/функций/операторов, но вариативные
функции позволяют достичь того же за счёт увеличения сложности кода.