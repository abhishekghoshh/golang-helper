package test

import (
	c "cards/card"
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
