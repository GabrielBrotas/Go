package main

import "fmt"

type Person struct {
	// address
	StreetAddress, PostCode, City string

	// job
	CompanyName, Position string
	AnnualIncome int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}


func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

// aggregators of the Person Builder -----

type PersonAddressBuilder struct {
	PersonBuilder
}

func (it *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = streetAddress
	return it
}
  
func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}
  
func (it *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	it.person.PostCode = postcode
	return it
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}
  
func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}
  
func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}


func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().At("Camacari BA").In("Brazil").WithPostcode("42800-000").
		Works().At("Wex").AsA("Software Engineer").Earning(999999)

	person := pb.Build()

	fmt.Println(person)

}