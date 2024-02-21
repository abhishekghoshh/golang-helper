package concurrency

import (
	"fmt"
	"time"
)

func evenSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i
		}
	}
	ch <- result
}
func squareSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i * i
		}
	}
	ch <- result
}

// https://gobyexample.com/select
func main() {
	evenCh := make(chan int)
	sqCh := make(chan int)

	go evenSum(0, 100, evenCh)
	go squareSum(0, 100, sqCh)

	select {
	case x := <-evenCh:
		fmt.Println(x)
	case y := <-sqCh:
		fmt.Println(y)
	default:
		fmt.Println("This is default case")
	}

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

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
