package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Testing the length of the deck
	if len(d) != 12 {
		t.Errorf("Expected 12 cards but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {	
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}
	if d[len(d) - 1] != "Three of Clubs" {
		t.Errorf("Expected last card to be Three of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 12 {
		t.Errorf("Expected 12 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}