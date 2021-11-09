package model

import "fmt"

var (
	// CompanyName содержит название компании
	CompanyName     = "test"
	companyLocation = "somecity"
)

// Структура Person
type Person struct {
	Name string
	age  int
}

// GetAge возвращает возраст человека
func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) getName() string {
	return p.Name
}

type company struct {
}

// GetPerson возвращает объект Person
func GetPerson() *Person {
	p := &Person{
		Name: "test",
		age:  21,
	}
	fmt.Println("Model Package:")
	fmt.Println(p.Name)
	fmt.Println(p.age)
	return p
}

func getCompanyName() string {
	return CompanyName
}
