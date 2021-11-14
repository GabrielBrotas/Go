package main

import "fmt"

// definition of a struct
type person struct {
	name    string
	age     int
	contact contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// ? create a struct ------------------------
	// assign value
	// firstPerson := person{
	// 	name: "Gabriel",
	// 	age: 21}

	// fmt.Println(firstPerson)

	// ? assign with dot notation --------------------------------------------------
	// var human person
	// human.name = "Gabriel"
	// human.age = 21
	// fmt.Println(human)

	// ? assign with struct literal ------------------------------------------------
	jim := person{
		name: "Jim",
		age:  25,
		contact: contactInfo{
			email:   "jim.gmail.com",
			zipCode: 94000,
		},
	}

	// print the whole value
	// fmt.Printf("%+v", jim)

	// ? update the name V1 ------------------------------------------------------------

	// jim.print()

	// &variable => give me the memory address value of this variable where is pointing at, is a cursor in memory
	// &jim => ex: 0xc00001 
	// jimPointer := &jim // assign the address of jim in memory to jimPointer

	// jimPointer.updateName("JonDoe")

	// jim.print()

	// Turn address into value with *address
	// Turn value into address with &value

	// ? update the name v2 ------------------------------------------------------------

	// this will work because the params of the function updateName says that it's a pointer to a person
	// so jim will be converted to a pointer automatically 
	// jim.updateName("Jondoe")
	// jim.print()
}

// print the person
func (p person) print() {
	fmt.Printf("%+v", p)
}


// func (p person) updateName(newName string) {
// *pointers => a pointer that points at a person, 
// this is a type description, it means we're working with a pointer to a person
func (pointerToPerson *person) updateName(newName string) {
	// this '*' means => Give me the value this memory address is pointing at 
	(*pointerToPerson).name = newName
}
