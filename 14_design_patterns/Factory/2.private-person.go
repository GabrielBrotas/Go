package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	Name string
	Age int
}

func(p *person) SayHello() {
	fmt.Println("Hii")
}

// that way the user will not be able to modify the person properties like name, age,...
func NewPerson(name string, age int) Person {
	return &person{name, age}
}

func main() {
	person := NewPerson("Gabriel", 22)
	person.SayHello()	
}
