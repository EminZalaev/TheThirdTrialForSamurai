package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

//Создаем сложные объекты по шагам, на каждом шаге происходит часть общего объекта
//выполняем шаги поочередно формируем объект, который представляет из себя сложную структуру
//используется один и тот же код строительства объекта для получения разных представлений этого объекта
//можем пропустить шаги или добавить еще шаг при создании объекта

//плюсы
//можно создать пошагово общий продукт который зависит от маленьких состовляющих частей
//позволяет использовать один и тот же код для создания различных объектов
//изолирует сложный код

//минусы
//усложняет код программы из за введений доп структур, интерфейсов
//клиент будет привязан к определенному объекту строителя, т.к интерфейс не может включать какой то метод и нужно будет включать его

import "fmt"

type Person struct {
	name, address, pin             string
	workAddress, company, position string
	salary                         int
}

type PersonBuilder struct {
	person *Person
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (a *PersonAddressBuilder) At(address string) *PersonAddressBuilder {
	a.person.address = address
	return a
}

func (a *PersonAddressBuilder) WithPostalCode(pin string) *PersonAddressBuilder {
	a.person.pin = pin
	return a
}

func (j *PersonJobBuilder) As(position string) *PersonJobBuilder {
	j.person.position = position
	return j
}

func (j *PersonJobBuilder) For(company string) *PersonJobBuilder {
	j.person.company = company
	return j
}

func (j *PersonJobBuilder) In(companyAddress string) *PersonJobBuilder {
	j.person.workAddress = companyAddress
	return j
}

func (j *PersonJobBuilder) WithSalary(salary int) *PersonJobBuilder {
	j.person.salary = salary
	return j
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func RunBuilderFacet() {
	pb := NewPersonBuilder()
	pb.Lives().At("Bangalore").
		WithPostalCode("560102").
		Works().
		As("Software Engineer").
		For("IBM").
		In("Bangalore").
		WithSalary(150000)

	person := pb.Build()

	fmt.Println(person)
}
