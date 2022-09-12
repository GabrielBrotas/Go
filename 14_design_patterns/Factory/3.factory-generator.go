package main

import "log"

type Employee struct {
	name, position, company string
	
}

// functional factory - high order function
func NewEmployeeFactory(position, company string) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, company}
	}
}

// class
type EmployeeFactory struct {
	Name, Position, Company string
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.Company}
}

func NewEmployeeFactory2(position, company string) *EmployeeFactory {
	return &EmployeeFactory{
		Position: position,
		Company: company,
	}
}

func main() {
	developersFactory := NewEmployeeFactory("developer", "Company A")

	developer1 := developersFactory("Ronald Wesley")
	developer2 := developersFactory("Fred Wesley")
	developer3 := developersFactory("George Wesley")

	log.Println(developer1)
	log.Println(developer2)
	log.Println(developer3)

	managersFactory := NewEmployeeFactory2("manager", "Comapany B")
	
	boss1 := managersFactory.Create("Arthur Wesley")
	boss2 := managersFactory.Create("Molly Wesley")

	log.Println(boss1)
	log.Println(boss2)
}