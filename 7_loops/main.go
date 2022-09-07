package main

import "fmt"

func main() {

	// for loop ------------
	for i:=0; i<=10; i++ {
		fmt.Println("Line = ", i)
	}

	// interact over array ------------
	var fruits []string

	fruits = append(fruits, "apple", "banana")

	for i, fruit := range fruits {
		fmt.Printf("Fruit %s in position %d \n", fruit, i)
	}

	// while ----
	var cars []string
	request_cars := []string{"Kwid", "Sandero", "Gol"}

	cars = append(cars, request_cars...)

	for len(request_cars) > 0 {
		request_cars = request_cars[:len(request_cars)-1] // remove the last index
		fmt.Printf("%v", request_cars)
	}
	fmt.Println()
	// interace over a string
	myString := "Hello World" // a string is a slice of bytes

	for i, letter := range myString {
		fmt.Println("Index =", i, "byte slice =", letter, "letter = ", string(letter))
	}

	// interact over an object
	var booksYear = map[string]string{"so": "2020"}

	booksYear["hp"] = "2000"
	booksYear["xpto"] = "2010"

	for key, year := range booksYear {
		fmt.Println("Book = ", key, "Year = ", year)
	}

}
