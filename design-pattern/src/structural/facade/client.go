package facade

import (
	"design-pattern/src/structural/facade/wallet"
	"fmt"
	"log"
)

func Main() {

	wallet := wallet.New("Abhishek Ghosh", 1234)
	fmt.Println()

	if err := wallet.Add("Abhishek Ghosh", 1234, 10); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()

	if err := wallet.Deduct("abc", 1234, 5); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
}
