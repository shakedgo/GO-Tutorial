package main

// import "fmt"

func main() {
	cards := newDeck()
	deckName := "MyDeck"
	// createDeckFile(deckName)
	cards.shuffle()
	createDeckFile(deckName, cards)
	newDeck := getNewDeck(deckName)
	newDeck.print()
	
	// hand, remainingCards := deal(cards, 5)
	// fmt.Println("Hand cards: ")
	// hand.print()
	// fmt.Println("Remaining cards: ")
	// remainingCards.print()
}

func createDeckFile(name string, d deck) {
	d.saveToFile(name)
}
func getNewDeck(name string) deck {
	return newDeckFromFile(name)
}