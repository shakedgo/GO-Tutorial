package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	fmt.Println("Hand cards: ")
	hand.print()
	fmt.Println("Remaining cards: ")
	remainingCards.print()
}