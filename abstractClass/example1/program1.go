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
