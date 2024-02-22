package concurrency

import (
	"fmt"
	"time"
)

// https://gobyexample.com/timeouts
func Timeouts() {

	c1 := make(chan string, 1)
	go func(ch chan string) {
		time.Sleep(2 * time.Second)
		ch <- "result 1"
	}(c1)

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func(ch chan string) {
		time.Sleep(2 * time.Second)
		ch <- "result 2"
	}(c2)
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
