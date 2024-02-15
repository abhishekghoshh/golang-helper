package flyweight

import "fmt"

func Main() {
	game := newGame()

	//Add Terrorist
	terroristDress, _ := getDressFactorySingleInstance().getDressByType(TerroristDressType)
	game.addTerrorist(TerroristDressType, terroristDress)
	game.addTerrorist(TerroristDressType, terroristDress)
	game.addTerrorist(TerroristDressType, terroristDress)
	game.addTerrorist(TerroristDressType, terroristDress)

	//Add CounterTerrorist
	counterTerrroristDress, _ := getDressFactorySingleInstance().getDressByType(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType, counterTerrroristDress)
	game.addCounterTerrorist(CounterTerrroristDressType, counterTerrroristDress)
	game.addCounterTerrorist(CounterTerrroristDressType, counterTerrroristDress)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
