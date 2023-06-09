package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

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

func (d deck) toString() string {
	// Convert deck to a slice of string
	sliceDeck := []string(d)
	return strings.Join(sliceDeck, ",")
}

func (d deck) saveToFile(name string) error {
	// WriteFlie(filename string, data []byte, perm os.FileMode) error
	// perm 0666 is anyone can read or write to the file
	return ioutil.WriteFile(name, []byte(d.toString()), 0666)
}
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}