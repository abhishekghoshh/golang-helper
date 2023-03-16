package main

import (
	"fmt"
)

// https://gobyexample.com/channel-synchronization
func worker(n int, done chan bool) {
	fmt.Print("working...")
	for i := 0; i < n; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(5, done)
	<-done
}
