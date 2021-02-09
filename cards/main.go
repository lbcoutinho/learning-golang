package main

import (
	"fmt"
	"os"
)

func main() {
	// Create new deck
	cards := newDeck()
	fmt.Println("==> New deck")
	cards.print()
	fmt.Println("\n========================")

	// Deal a head with 5 cards
	hand, remainingCards := deal(cards,5)
	fmt.Println("==> Deal")
	hand.print()
	fmt.Println("------------------------")
	fmt.Println("==> Remaining cards")
	remainingCards.print()
	fmt.Println("\n========================")

	// Shuffle the deck
	fmt.Println("==> Shuffled deck")
	cards.shuffle()
	cards.print()
	fmt.Println("\n========================")

	// Save the deck to a file
	filename := "cards"
	fmt.Println("==> Saving deck to file...")
	err := cards.saveToFile(filename)
	handleError(err)
	fmt.Println("==> Saved successfully to file", filename)
	fmt.Println("\n========================")

	// Load the deck from a file
	fmt.Println("==> Loading deck from file...")
	d, err := newDeckFromFile(filename)
	handleError(err)
	fmt.Println("==> Loaded file", filename, "successfully")
	d.print()
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
