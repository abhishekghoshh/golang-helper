package card

import (
	c "basics/testings/src/card"
	"testing"
)

func TestNewCards(t *testing.T) {
	cards := c.NewCards()
	l := len(cards)
	if l != 32 {
		t.Errorf("Expected deck length of 32, but got %v", l)
	}
	t.Log("TestNewCards passed")
}

func TestGiveDeck(t *testing.T) {
	cards := c.NewCards()
	firstDeck, secondDeck := cards.GiveDeck()
	if len(firstDeck) != 4 || len(secondDeck) != 4 {
		t.Errorf("Expected deck length of 16 and 16, but got %v %v", len(firstDeck), len(secondDeck))
	}
	t.Log("TestGiveDeck passed")
}
