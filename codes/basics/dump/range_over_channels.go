package concurrency

import "fmt"

// https://gobyexample.com/range-over-channels
func RangeOverChannels() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	fmt.Println("length of channel", len(queue))
	for elem := range queue {
		fmt.Println("length of channel", len(queue), ", channel content", elem)
	}
}