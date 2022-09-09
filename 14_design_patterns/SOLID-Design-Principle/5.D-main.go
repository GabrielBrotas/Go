package main

import (
	"log"
)

/*
	Dependency Inversion Principle
	High Level Model (HLM) shoud not depend on Low Level Model (LLM)
	Both should depend on abstractions
*/

type Relashionship int

const (
	Parent Relashionship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person       // Jon
	relationship Relashionship // sibling of
	to           *Person       // Arthur
}

// Low Level Model -> a database persistence
type RelationshipBrowser interface {
	FindAllChildrenOf(parent_name string) []*Person
}

type Relashionships struct {
	relations []Info
}

func (r *Relashionships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
}

// implement the interface
func (r *Relashionships) FindAllChildrenOf(parent_name string) []*Person {
	var children []*Person
	
	for i, person := range r.relations {
		if person.from.name == parent_name && person.relationship == Parent {
			children = append(children, r.relations[i].to)
		}
	}

	return children
}

// High Level Model -> Model to operate in the data model(LLM)
type Research struct {
	// ! Break DIP
	// ! relationships Relashionships // this is gonna work, but we are depending on that particular model to work

	// ? DIP
	browser RelationshipBrowser
}

func (r *Research) Investigate(parent *Person) {
	// ! Break DIP
	// if the relationships model decide to change its mechanism, instead of in memory be a database
	// this code will break, becuase it depends directly from another model 
	// relations := r.relationships.relations
	// for _, rel := range relations {
	// 	if rel.from.name == "Arthur" && rel.relationship == Parent {
	// 		log.Println("Arthur has a child called", rel.to.name)
	// 	}
	// }

	// ? DIP
	children := r.browser.FindAllChildrenOf(parent.name)

	for _, child := range children {
		log.Println("Arthur has a child called", child.name)
	}

}

func main() {
	father := Person{"Arthur"}
	son1 := Person{"Ronald"}
	son2 := Person{"Fred"}
	son3 := Person{"George"}

	relationships := Relashionships{}
	relationships.AddParentAndChild(&father, &son1)
	relationships.AddParentAndChild(&father, &son2)
	relationships.AddParentAndChild(&father, &son3)
	
	research := Research{&relationships}
	research.Investigate(&father)
}
