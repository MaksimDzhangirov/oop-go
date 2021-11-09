package view

import (
	"fmt"
	"github.com/MaksimDzhangirov/OOP-go/encapsulation/model"
)

// функция Test
func Test() {
	// ИДЕНТИФИКАТОР СТРУКТУРЫ
	p := &model.Person{
		Name: "test",
		age:  21,
	}
	fmt.Println(p)
	c := &model.company{}
	fmt.Println(c)

	// МЕТОД СТРУКТУРЫ
	fmt.Println(p.GetAge())
	fmt.Println(p.getName())

	// ПОЛЯ СТРУКТУРЫ
	fmt.Println(p.Name)
	fmt.Println(p.age)

	// ФУНКЦИЯ
	person2 := model.GetPerson()
	fmt.Println(person2)
	companyName := model.getCompanyName()
	fmt.Println(companyName)

	// ПЕРЕМЕННЫЕ
	fmt.Println(model.companyLocation)
	fmt.Println(model.CompanyName)
}