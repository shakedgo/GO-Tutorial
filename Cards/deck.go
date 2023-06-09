package main

import "fmt"

type deck []string

// Create a new deck function
func newDeck() deck {
	cards := deck{}
	
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}
// A function which returns two values of type deck
func deal(d deck, handSize int) (deck, deck) {
	// return deck from 0 to handSize, and the remaining cards in the deck (from handSize to end)
	return d[:handSize], d[handSize:]
}

// Reciver = kinda like deck is a class and print is a method of that class
func (d deck) print() {
	for i, cards := range d {
		fmt.Println(i, cards)
	}
}