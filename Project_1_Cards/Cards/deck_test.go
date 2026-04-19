package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// Refactor: put the file name in a const to avoid typos
	const FILE = "_decktesting"

	// Clean up: delete any files with the name '_decktesting'
	os.Remove(FILE)

	// Create a new deck and save it to the Hard Drive
	d := newDeck()
	d.saveToFile(FILE)

	// Load a new deck from the file
	ld := newDeckFromFile(FILE)

	// Assert that the loaded deck length and deck length are equal
	if len(ld) != len(d) {
		t.Errorf("Expected %v cards in deck, got %v", len(d), len(ld))
	}

	// Refactor: Assert that the loaded deck is the same as deck
	if ld.toString() != d.toString() {
		t.Errorf("Deck loaded is not the same as the saved one. got: %v, want: %v", d, ld)
	}

	// Clean up: delete any files with the name '_decktesting'
	os.Remove(FILE)
}
