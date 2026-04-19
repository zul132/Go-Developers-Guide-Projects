package main

/*
	Go is NOT an object-oriented programming language
	==> there is no concept of "classes" in Go
*/

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func main() {
// 	cards := newDeckFromFile("my_cards")
// 	cards.print()

// 	// cards := newDeckFromFile("my") --> this will throw error: no such file or directory
// }

// cards := newDeck()

// hand, remainingCards := deal(cards, 5)

// hand.print()
// remainingCards.print()
