package main

import "fmt"

type Person struct {
	// personal
	name string
	age int

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

func (b *PersonBuilder) Personal() *PersonPersonalBuilder {
	return &PersonPersonalBuilder{*b}
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

type PersonPersonalBuilder struct {
	PersonBuilder
}

func (ppb *PersonPersonalBuilder) Called(name string) *PersonPersonalBuilder {
	ppb.person.name = name
	return ppb
}

func (ppb *PersonPersonalBuilder) Age(age int) *PersonPersonalBuilder {
	ppb.person.age = age
	return ppb
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Personal().Called("Gabriel").Age(22).
		Lives().At("Camacari BA").In("Brazil").WithPostcode("00000-000").
		Works().At("Company A").AsA("Software Engineer")

	person := pb.Build()

	fmt.Println(person)
}