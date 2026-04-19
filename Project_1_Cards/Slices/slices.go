package main

import "fmt"

/*
	Go has two basic data structures for handling lists of records :-
		1. Array - Fixed length list of records
		2. Slice - An array that can grow or shrink

	Every element in an array or slice must be of the same type
*/

func main() {
	// create a slice
	cards := []string{"Ace of Spades", newCard()}

	// appending to a slice - add a new element at the end of a slice
	/* The append function doesn't modify the existing slice,
	   instead it returns a new slice */
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
