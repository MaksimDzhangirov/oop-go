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
	base  // встраиваем
	style string
}

func check(b iBase) {
	b.say()
}

func main() {
	base := base{
		color: "Red",
		clear: func() {
			fmt.Println("Clear from child's function")
		},
	}
	child := &child{
		base:  base,
		style: "somestyle",
	}
	child.say()
}
