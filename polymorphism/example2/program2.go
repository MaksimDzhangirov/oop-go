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
