package concurrency

import (
	"fmt"
	"time"
)

// https://gobyexample.com/channels
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
