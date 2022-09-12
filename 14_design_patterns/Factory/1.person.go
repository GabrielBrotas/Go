package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

func main() {
	person := NewPerson("Gabriel", 22)

	fmt.Println(person)

}
