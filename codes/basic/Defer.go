package main

import (
	"fmt"
)

func run() {
	fmt.Println("run method is executed")
}

// https://gobyexample.com/defer
func main() {
	defer run()
	fmt.Println("The main method starts")
	func() {
		defer run()
		fmt.Println("In the anonymous function")
	}()
	fmt.Println("The main method finished")
}
