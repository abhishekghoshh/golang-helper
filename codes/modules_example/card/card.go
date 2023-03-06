package card

import "fmt"

// Create a new type of 'Card'
// which is a slice of strings
type Cards []string

func NewCards() Cards {
	cards := Cards{}
	cardType := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Jack", "Nine", "Ace", "Ten", "King", "Queen", "Eight", "Seven"}
	for _, cardType := range cardType {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+cardType)
		}
	}
	return cards
}

func (cards *Cards) Print() {
	for _, val := range *cards {
		fmt.Println(val)
	}
}
