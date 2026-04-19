package main

import "fmt"

// We can initialize a variable outside the main func, but we can only assign a value to itside main
var deckSize int

func main() {
	// var card string = "Ace of Spades"

	// We only use the := operator when initializing a NEW variable
	card := "Ace of Spades"

	// Don't include colon when assigning a new value to a variable
	card = "Five of Diamonds"

	// card = 5 will not worked as we had initialized "card" with a string value,
	// so Go will infer that the variable type is String from the initialized value

	fmt.Println(card)

	deckSize = 52
	fmt.Println(deckSize)

	// Variables must first be initialized with either the := operator or the "var variableName type" syntax
}
