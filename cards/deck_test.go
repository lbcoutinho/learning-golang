package main

import (
	"os"
	"testing"
)

const testFile = "_decktesting"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	const expectedLength = 52
	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %d, but got %d", expectedLength, len(d))
	}

	const expectedFirstCard = "Ace of Spades"
	if d[0] != expectedFirstCard {
		t.Errorf("Expected deck first card equals %s, but got %s", expectedFirstCard, d[0])
	}

	const expectedLastCard = "King of Clubs"
	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected deck last card equals %s, but got %s", expectedLastCard, d[len(d)-1])
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	const handsize = 5
	dealResult, remainingCards := deal(d, handsize)

	if len(dealResult) != 5 {
		t.Errorf("Expected deal of %d cards, but got %d", handsize, len(dealResult))
	}

	expectedRemainingCards := len(d) - handsize
	if len(remainingCards) != expectedRemainingCards {
		t.Errorf("Expected deal of %d cards, but got %d", expectedRemainingCards, len(remainingCards))
	}
}

func TestToString(t *testing.T) {
	d := newDeck()
	cards, _ := deal(d, 3)
	result := cards.toString()

	expectedResult := cards[0] + "," + cards[1] + "," + cards[2]
	if result != expectedResult {
		t.Errorf("Expected toString of '%s', but got '%s'", expectedResult, result)
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	d := newDeck()
	os.Remove(testFile)
	d.saveToFile(testFile)
	readDeck, _ := newDeckFromFile(testFile)

	if len(d) != len(readDeck) {
		t.Errorf("Expected save deck to have length of %d, but got %d", len(d), len(readDeck))
	}

	os.Remove(testFile)
}
