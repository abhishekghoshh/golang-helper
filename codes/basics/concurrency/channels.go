package concurrency

import (
	"fmt"
	"time"
)

/*
https://gobyexample.com/channels

Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and receive those values into another goroutine.

Create a new channel with make(chan val-type). Channels are typed by the values they convey.

	messages := make(chan string)

Send a value into a channel using the channel <- syntax.
Here we send "ping" to the messages channel we made above, from a new goroutine.

	go func() { messages <- "ping" }()

The <-channel syntax receives a value from the channel. Here we'll receive the "ping" message we sent above and print it out.

	msg := <-messages
	fmt.Println(msg)

When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.
By default sends and receives block until both the sender and receiver are ready.
This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
*/
func Channels() {
	var out = func(from, to int, ch chan bool) {
		for i := from; i <= to; i++ {
			time.Sleep(50 * time.Millisecond)
			fmt.Println(i)
		}
		ch <- true
	}

	var evenSum = func(from, to int, ch chan int) {
		result := 0
		for i := from; i <= to; i++ {
			if i%2 == 0 {
				result += i
			}
		}
		ch <- result
	}
	var squareSum = func(from, to int, ch chan int) {
		result := 0
		for i := from; i <= to; i++ {
			if i%2 == 0 {
				result += i * i
			}
		}
		ch <- result
	}

	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go out(0, 5, ch1)
	go out(6, 10, ch2)

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)

	evenCh := make(chan int)
	sqCh := make(chan int)

	go evenSum(0, 100, evenCh)
	go squareSum(0, 100, sqCh)

	fmt.Println(<-evenCh + <-sqCh)
}
