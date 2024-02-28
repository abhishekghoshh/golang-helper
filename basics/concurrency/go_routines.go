package concurrency

import (
	"fmt"
	"time"
)

/*
https://gobyexample.com/goroutines

A goroutine is a lightweight thread of execution.
*/

func GoRoutines() {
	var out = func(from, to int) {
		for i := from; i <= to; i++ {
			fmt.Print(i, " ")
		}
	}

	go out(0, 5)
	go out(6, 10)
	time.Sleep(time.Second)
	fmt.Println()

	// Suppose we have a function call f(s). Here's how we'd call that in the usual way, running it synchronously.

	var f = func(from string) {
		for i := 0; i < 3; i++ {
			fmt.Println(from, ":", i)
			time.Sleep(100 * time.Millisecond)
		}
	}
	// To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.
	go f("goroutine 1")

	f("direct")

	go f("goroutine 2")

	// You can also start a goroutine for an anonymous function call.
	go func(msg string) {
		fmt.Println(msg)
		time.Sleep(100 * time.Millisecond)
	}("go routine 3 single message")

	// Our two function calls are running asynchronously in separate goroutines now.
	// Wait for them to finish (for a more robust approach, use a WaitGroup).
	time.Sleep(2 * time.Second)

	// When we run this program, we see the output of the blocking call first, then the output of the two goroutines.
	// The goroutines' output may be interleaved, because goroutines are being run concurrently by the Go runtime.
	fmt.Println("done processing")
}
