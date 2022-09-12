package main

import "fmt"

type Person struct {
	name, position string
}

type personModifications func(*Person) // modifications applied to this person
type PersonBuilder struct {
	actions []personModifications
}

// instead of commit already we can create small transactions and build at the end
func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

// transactions
func (b *PersonBuilder) WorkAs(position string) *PersonBuilder {
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func main() {
	b := PersonBuilder{}
	// we 
	person := b.Called("Gabriel").Build()

	fmt.Println(person)
}