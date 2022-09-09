package main

import "fmt"

/*
	Open Closed Principle

	open for extension but closed for modification

	Specification - Enterprise Pattern
*/

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	// properties...
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) FilterBySyze(products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) FilterBySyzeAndColor(products []Product, size Size) *[]Product {
	return &products
}

// This is a violation of the open closed principle, we should be able to extend a scenario based on conditions, types, rules...
// without modifying something you've already written
// we should have one Filter func and this method should be open for extension and return the same struct
// extendable setup

// Specification Pattern ---
type Specification interface {
	isSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) isSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) isSatisfied(p *Product) bool {
	return p.size == s.size
}

type CompositionSpecification struct {
	first, second Specification
}

func (s CompositionSpecification) isSatisfied(p *Product) bool {
	return s.first.isSatisfied(p) && s.second.isSatisfied(p)
}

type BetterFilter struct {}

// we dont have to modify this Filter anymore, the only thing we may add is more specification
// and these specification just have to follow the Specification interface
func (b *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		// now we have a more flexible filter because the specification will define the filter type
		if spec.isSatisfied(&v) { 
			result = append(result, &products[i])
		}
	}

	return result
}


func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products (old):\n")

	f := Filter{}
	
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	for _, v := range f.FilterBySyze(products, large) {
		fmt.Printf(" - %s is large\n", v.name)
	}

	// ----- with better filter and specification pattern ---------- 
	betterFilter := BetterFilter{}
	colorSpec := ColorSpecification{blue}
	
	fmt.Printf("Blue products (new):\n")

	for _, v := range betterFilter.Filter(products, colorSpec) {
		fmt.Printf(" - %s better filter is blue\n", v.name)
	}

}
