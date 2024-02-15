package flyweight

type game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (c *game) addTerrorist(dressType string, dress Dress) {
	player := newPlayer("T", dressType, dress)
	c.terrorists = append(c.terrorists, player)
}

func (c *game) addCounterTerrorist(dressType string, dress Dress) {
	player := newPlayer("CT", dressType, dress)
	c.counterTerrorists = append(c.counterTerrorists, player)
}
