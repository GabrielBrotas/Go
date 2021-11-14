package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// Create a new function that returns a value of type 'deck'
// d is a variable of type 'deck' that we created above,
// this function is called when we use the variable of type deck .print(), d is the value of deck
// any variable of type deck can be used as a parameter of this function
// receiver is the variable of type deck that we created above
// we can think in 'd' as 'this' or 'self'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//  -----------------------

func newDeck() deck {
	cards := deck{}

	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}	

	return cards
}

//  -----------------------

// return two separete decks
func deal(d deck, handSize int) (deck, deck) {
	// return everything from the start of the deck to the handSize,
	// and everything from handSize to the end of the deck
	// is like split the deck in two
	return d[:handSize], d[handSize:]
}

// assign the deal function to the deck type
// func (d deck) deal(handSize int) (deck, deck) {
// 	return d[:handSize], d[handSize:]
// }

//  -----------------------

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// if a error occur we can return it
// this parameters we need to use in the WriteFile function from ioutil
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // 0666 is the permission of the file that anyone can read and write
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		// Option #1 - Log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error: ", err)
		// operate system exit, 0 is successful and non-zero represents an error
		os.Exit(1)
	}

	s := strings.Split(string(byteSlice), ",") // Ace of Spades,Two of Spades, ....

	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // 	create a new source of random numbers
	r := rand.New(source) // create a new random number generator
	
	for i := range d {
		// generate a random number from 0 to the length of the deck
		newPosition := r.Intn(len(d) - 1)

		// swap the card at position i with the card at position newPosition
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}