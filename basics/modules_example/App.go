package main

import (
	c "cards/card"
	"fmt"
)

func main() {
	cards := c.NewCards()
	cards.Shuffle()
	// cards.Print()
	//fmt.Println(len(cards))

	fmt.Println(cards.ToString())

	firstDeck, secondDeck := cards.GiveDeck()
	fmt.Println(firstDeck, secondDeck)

	cards.SaveToFile("allcards.txt")
	fmt.Println(c.NewCardsFromFile("allcards.txt").ToString())
}
