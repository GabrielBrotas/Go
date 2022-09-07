package main

import "fmt"

type bot interface {
	getGreeting() string
}


type englishBot struct{}
type portuguesBot struct{}


func main() {
	eb := englishBot{}
	pt := portuguesBot{}

	printGreeting(eb)
	printGreeting(pt)
}


func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/*
 that means you implement bot interface
*/
func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (portuguesBot) getGreeting() string {
	return "Oi!"
}
// ! BAD CODE ------------------------------------------------------------
// func main() {
	// eb := englishBot{}
	// pt := portuguesBot{}

	// fmt.Println(eb.getGreeting())
	// fmt.Println(pt.getGreeting())
	// printGreeting(eb)
// }


// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func (englishBot) getGreeting() string {
// 	return "Hi there!"
// }

// func (portuguesBot) getGreeting() string {
// 	return "Oi!"
// }