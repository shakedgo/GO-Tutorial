package main

// import "fmt"

func main() {
	cards := newDeck()
	
	createDeckFile(cards)
	newDeck := getNewDeck()
	newDeck.print()
	
	// hand, remainingCards := deal(cards, 5)
	// fmt.Println("Hand cards: ")
	// hand.print()
	// fmt.Println("Remaining cards: ")
	// remainingCards.print()
}

func createDeckFile(d deck) {
	d.saveToFile("MyDeck")
}
func getNewDeck() deck {
	return newDeckFromFile("MyDeck")
}