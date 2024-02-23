package concurrency

import (
	"fmt"
	"time"
)

/*
https://gobyexample.com/select

Go's select lets you wait on multiple channel operations.
Combining goroutines and channels with select is a powerful feature of Go.

For our example we'll select across two channels.
*/
func Selects() {
	// Each channel will receive a value after some amount of time,
	// to simulate e.g. blocking RPC operations executing in concurrent goroutines.
	c1 := make(chan string)
	c2 := make(chan string)

	go func(ch chan string) {
		time.Sleep(1 * time.Second)
		ch <- "one"
	}(c1)

	go func(ch chan string) {
		time.Sleep(2 * time.Second)
		ch <- "two"
	}(c2)

	// We'll use select to await both of these values simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	// We receive the values "one" and then "two" as expected.
	// Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.

	// now we will use a even sum and square sum function and go routine in them
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
	evenCh := make(chan int)
	sqCh := make(chan int)

	go evenSum(0, 100, evenCh)
	go squareSum(0, 100, sqCh)

	// if none of the channels are ready then it will go to the default case
	select {
	case x := <-evenCh:
		fmt.Println(x)
	case y := <-sqCh:
		fmt.Println(y)
	default:
		fmt.Println("This is default case")
	}
}
