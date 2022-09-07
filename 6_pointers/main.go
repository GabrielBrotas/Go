package main

import "fmt"

func main() {
	var car string = "Renault"

	fmt.Println(car)

	changeMemoryAddressValue(&car)

	fmt.Println(car)
}

func changeMemoryAddressValue(addrRef *string) {
	fmt.Println(addrRef) // memory address

	*addrRef = "Ford"
}