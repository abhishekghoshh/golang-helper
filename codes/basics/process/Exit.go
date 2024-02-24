package process

import (
	"fmt"
	"os"
)

// https://gobyexample.com/exit
func DoExit() {
	defer fmt.Println("!")
	os.Exit(3)
}
