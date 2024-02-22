package concurrency

import (
	"fmt"
	"time"
)

// https://gobyexample.com/goroutines
func GoRoutines() {
	var out = func(from, to int) {
		for i := from; i <= to; i++ {
			fmt.Print(i)
		}
	}
	var f = func(from string) {
		for i := 0; i < 3; i++ {
			fmt.Println(from, ":", i)
			time.Sleep(10 * time.Millisecond)
		}
	}

	go out(0, 5)
	go out(6, 10)
	time.Sleep(time.Second)
	fmt.Println()
	go f("goroutine")
	f("direct")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(5 * time.Second)
	fmt.Println("done")
}
