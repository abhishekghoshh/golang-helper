package main

import (
	"fmt"
	"time"
)

// https://gobyexample.com/tickers
func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func(ticker *time.Ticker) {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}(ticker)

	time.Sleep(20000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
