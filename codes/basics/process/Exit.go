package main

import (
	"fmt"
	"os"
)

// https://gobyexample.com/exit
func main() {
	defer fmt.Println("!")
	os.Exit(3)
}
