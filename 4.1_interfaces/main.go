package main

import "fmt"

type Animal interface {
	Says() string
}

type Dog struct {
	Name  string
	Color string
}

type Cat struct {
	Name string
	Legs int
}

func main() {
	dog := Dog{Name: "Dog X", Color: "Blue"}

	sayHi(&dog)
}

func sayHi(a Animal) {
	fmt.Println(a.Says())
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Cat) Says() string {
	return "Woof"
}
