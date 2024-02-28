package introduction

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var deckSize int

func HellowWorld() {
	deckSize = 50
	fmt.Println(deckSize)

	fmt.Println("Hello World")

	var input int = 10
	// fmt.Scanln(&input)

	fmt.Println("input is", input)

	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favourite number is", rand.Intn(100))

	fmt.Printf("Sqrt of 7 is %g \n", math.Sqrt(7))

	fmt.Println("value of PI is", math.Pi)

}
