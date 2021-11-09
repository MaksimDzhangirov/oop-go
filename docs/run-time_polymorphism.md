# Полиморфизм во время выполнения в Go

[Оригинал](https://golangbyexample.com/runtime-polymorphism-go/)

Полиморфизм во время выполнения означает, что решение о том какую функцию вызывать 
принимается во время выполнения.

Давайте разберём этот момент на примере. В разных странах по-разному рассчитывается
налог. Это поведение можно представить с помощью интерфейса.

```go
type taxCalculator interface {
	calculateTax()
}
```

Теперь у разных стран может быть различная структура, но все они будут реализовывать 
метод `calculateTax()`. Например, структура `indianTax` показана ниже. В ней также
может быть определен метод `calculateTax()`, который будет выполнять реальные 
вычисления, используя проценты.

Точно так же налоговые системы других стран могут быть представлены структурой,
и они также будут реализовывать свой собственный метод `calculateTax()` для
определения значения налога.

Теперь давайте посмотрим, как мы можем использовать этот интерфейс `taxCalcuator`
для расчета налога с лиц, проживающих в разных странах в различное время года. 
Код всей программы показан ниже:

```go
package main

import "fmt"

type taxSystem interface {
    calculateTax() int
}

type indianTax struct {
    taxPercentage int
    income        int
}

func (i *indianTax) calculateTax() int {
    tax := i.income * i.taxPercentage / 100
    return tax
}

type singaporeTax struct {
    taxPercentage int
    income        int
}

func (i *singaporeTax) calculateTax() int {
    tax := i.income * i.taxPercentage / 100
    return tax
}

type usaTax struct {
    taxPercentage int
    income        int
}

func (i *usaTax) calculateTax() int {
    tax := i.income * i.taxPercentage / 100
    return tax
}

func main() {
    indianTax := &indianTax{
        taxPercentage: 30,
        income:        1000,
    }
    singaporeTax := &singaporeTax{
        taxPercentage: 10,
        income:        2000,
    }
    taxSystems := []taxSystem{indianTax, singaporeTax}
    totalTax := calculateTotalTax(taxSystems)
    
    fmt.Printf("Total Tax is %d\n", totalTax)
}

func calculateTotalTax(taxSystems []taxSystem) int {
    totalTax := 0
    for _, t := range taxSystems {
        totalTax += t.calculateTax() // вот где происходит полиморфизм во время выполнения
    }
    return totalTax
}
```

Результат в терминале:

```shell
go run polymorphism/example3/program3.go
Total Tax is 500
```

Ниже приведена строка, где происходит полиморфизм во время выполнения.

```go
totalTax += t.calculateTax() // вот где происходит полиморфизм во время выполнения
```

Один и тот же метод calculateTax используется в различных контекстах для расчёта
налога. Когда компилятор видит такой вызов, он откладывает решение о том какой 
метод будет вызван до времени выполнения. Это то, что происходит под капотом.

## Добавляем налоговые системы:

Теперь давайте расширим вышеприведенную программу, включив в нее налоговую 
систему для США.

```go
type usaTax struct {
    taxPercentage int
    income        int
}

func (i *usaTax) calculateTax() int {
    tax := i.income * i.taxPercentage / 100
    return tax
}
```

Нам достаточно изменить функцию main, чтобы добавить налоговую систему США.

```go
func main() {
    indianTax := &indianTax{
        taxPercentage: 30,
        income:        1000,
    }
    singaporeTax := &singaporeTax{
        taxPercentage: 10,
        income:        2000,
    }
    usaTax := &usaTax{
        taxPercentage: 40,
        income:        500,
    }
    taxSystems := []taxSystem{indianTax, singaporeTax, usaTax}
    totalTax := calculateTotalTax(taxSystems)
    
    fmt.Printf("Total Tax is %d\n", totalTax)
}
```

Результат в терминале:

```shell
go run polymorphism/example3/program3.go
Total Tax is 700
```

Как вы возможно заметили изменять функцию calculateTotalTax в вышеприведенной 
программе не нужно, чтобы учесть налоговую систему США. В этом заключается 
преимущество использования интерфейсов и полиморфизма.