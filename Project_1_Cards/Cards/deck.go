package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Create a new Deck of playing cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// '_' indicates that there is a variable, but you don't want to use it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Create a 'print' func with a Receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Go supports returning multiple values from one function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Helper function to convert a Deck into a String
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// ioutil package is Depecrated, so we'll use os package instead
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// os.ReadFile will return a slice of bytes from the read file
	bs, err := os.ReadFile(filename)

	// A successful call returns err == nil
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Convert the slice of bytes into a string
	// Use strings.Split() to split the converted string into a slice of playing cards
	s := strings.Split(string(bs), ",")

	return deck(s)
}

// Shuffle the order of cards in the deck using a randomized swapping logic
func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i int, j int) {
		d[i], d[j] = d[j], d[i]
	})

	// for i := range d {
	//     newPosition := rand.Intn(len(d) - 1)

	//     d[i], d[newPosition] = d[newPosition], d[i]
	// }
}

/*
func (d deck) shuffle() {
	// Iterate over each index in the deck
	for i := range d {
		// Note: No need to seed the rand with a random value from Go 1.20 onwards.
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)

		// generate a random value between 0 and len(d) - 1
		newPosition := r.Intn(len(d))

		// Simpler way to seed the rand (deprecated)
		rand.Seed(time.Now().UnixNano())
		newPosition := rand.Intn(len(d))

		// swap the elements at i and newPosition of the slice
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
*/
