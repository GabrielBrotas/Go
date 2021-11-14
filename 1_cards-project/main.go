package main

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	// ? basic variables declaration, function call and array interaction
	// var card string = "Ace of spades"
	// := create a variable, reassing a value we just use '='
	// card := "Ace of spades"
	// card = "new value"
	// card := newCard()

	// cards := []string{newCard(), "Two of Diamonds"}
	// cards := deck{newCard(), "Two of Diamonds"}
	// cards = append(cards, "Six of Diamonds") // add to the end of the slice the value

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards.print()

	// fmt.Println(card, cards)	

	// ? Functions call, multiple params ---------------------

	// cards := newDeck()

	// cards.print()
	// hand, remainingCards := cards.deal(5)

	// hand.print()
	// fmt.Println("--------")
	// remainingCards.print()

	// ? String to Byte ---------------------
	// byte is a ascii character representation of a string
	// greeting := "Hi bro"
	// fmt.Println([]byte(greeting)) // [72 105 32 98 114 111]

	// ? Write in the hard drive ---------------------
	
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// ? Read from the hard drive ---------------------
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// ? Shuffle ---------------------
	cards := newDeck()
	cards.shuffle()

	cards.print()

}
