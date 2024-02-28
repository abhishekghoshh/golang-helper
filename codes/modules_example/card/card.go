package card

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'Card' which is a slice of strings
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
func (c *Cards) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	cards := *c
	for i := range cards {
		newPosition := r.Intn(len(cards) - 1)
		cards[i], cards[newPosition] = cards[newPosition], cards[i]
	}
}
func (c *Cards) GiveDeck() ([][]string, [][]string) {
	c.Shuffle()
	allCards := []string(*c)
	firstDeck := make([][]string, 4)
	secondDeck := make([][]string, 4)
	if len(allCards) == 0 {
		return firstDeck, secondDeck
	}
	for i := 0; i < 4; i++ {
		firstDeck[i] = allCards[i*4 : ((i+1)*4)-1]
		secondDeck[i] = allCards[16+i*4 : 16+((i+1)*4)-1]
	}
	return firstDeck, secondDeck
}
func (cards *Cards) Print() {
	for _, val := range *cards {
		fmt.Println(val)
	}
}
func (cards Cards) ToString() string {
	return strings.Join([]string(cards), ",")
}
func (cards Cards) SaveToFile(filename string) error {
	return os.WriteFile(filename, []byte(cards.ToString()), 0666)
}

func NewCardsFromFile(filename string) Cards {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return Cards(s)
}
