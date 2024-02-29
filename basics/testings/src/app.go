package src

import (
	c "basics/testings/src/card"
	"fmt"
)

func SayHello() string {
	return "hello"
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func SavePlayingCards() {
	cards := c.NewCards()
	cards.Shuffle()
	firstDeck, secondDeck := cards.GiveDeck()
	fmt.Println(firstDeck, secondDeck)
	cards.SaveToFile("cards.txt")
}
