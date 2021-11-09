# Инкапсуляция в Go

[Оригинал](https://golangbyexample.com/encapsulation-in-go/)

Golang обеспечивает инкапсуляцию на уровне пакета. В Go нет ключевых слов 
`public`, `private` или `protected`. Единственный механизм управления 
видимостью — использование прописных и строчных букв.

* Идентификаторы с прописной буквы экспортируются. Прописная буква означает, 
  что это экспортируемый идентификатор.
* Идентификаторы со строчной буквы не экспортируются. Строчные буквы 
  указывают на то, что идентификатор не экспортируется и будет доступен 
  только из того же пакета.
  
Существует пять видов идентификаторов, которые можно экспортировать.

1. Структура
2. Метод структуры
3. Поле структуры
4. Функция
5. Переменная

Посмотрим на пример, показывающий как экспортируются все вышеперечисленные 
идентификаторы. Файл называется `data.go`, а пакет - `model`.

* Структура
    * Структура `Person` экспортируется
    * Структура `company` не экспортируется
* Метод структуры
    * Метод структуры `Person` `GetAge()` экспортируется
    * Метод структуры `Person` `getName()` не экспортируется
* Поле структуры
    * Поле структуры `Person` `Name` экспортируется
    * Поле структуры `Person` `age` не экспортируется
* Функция
    * Функция `GetPerson()` экспортируется
    * Функция `getCompanyName()` не экспортируется
* Переменные
    * Переменная `CompanyName` экспортируется
    * Переменная `companyLocation` не экспортируется
    
**data.go**

```go
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
```

Давайте добавим файл `data_test.go` в пакет `model`. Ниже показано содержимое файла.

**data_test.go**
```go
package model

import "fmt"

// функция Test
func Test() {
    // ИДЕНТИФИКАТОР СТРУКТУРЫ
    p := &Person{
        Name: "test",
        age:  21,
    }
    fmt.Println(p)
    c := &company{}
    fmt.Println(c)
    
    // МЕТОД СТРУКТУРЫ
    fmt.Println(p.GetAge())
    fmt.Println(p.getName())
    
    // ПОЛЯ СТРУКТУРЫ
    fmt.Println(p.Name)
    fmt.Println(p.age)
    
    // ФУНКЦИЯ
    person2 := GetPerson()
    fmt.Println(person2)
    companyName := getCompanyName()
    fmt.Println(companyName)
    
    // ПЕРЕМЕННЫЕ
    fmt.Println(companyLocation)
    fmt.Println(CompanyName)
}
```

При запуске этого файла он может получить доступ ко всем экспортированным и 
не экспортированным полям в `data.go`, поскольку оба находятся в одном и той же
пакете `model`. Ошибок при компиляции не возникло и мы получаем следующий 
результат.

Результат в терминале:

```shell
go run main.go
&{test 21}
&{}
21
test
test
21
Model Package:
test
21
&{test 21}
test
somecity
test
```

Давайте переместим файл `data_test.go` в другой пакет `view`. Теперь обратите
внимание на результат выполнения команды `go build`. Получаем ошибки компиляции.
Все ошибки компиляции связаны с невозможностью ссылаться на неэкспортированные 
поля.

```go
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
```

Результат в терминале:

```shell
go build .
view/test.go:13:3: cannot refer to unexported field 'age' in struct literal of type model.Person
view/test.go:16:8: cannot refer to unexported name model.company
view/test.go:21:15: p.getName undefined (cannot refer to unexported field or method model.(*Person).getName)
view/test.go:25:15: p.age undefined (cannot refer to unexported field or method age)
view/test.go:30:17: cannot refer to unexported name model.getCompanyName
view/test.go:34:14: cannot refer to unexported name model.companyLocation
```