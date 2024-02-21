package concurrency

import "fmt"

func addMessage(messages chan string) {
	messages <- "buffered"
	messages <- "channel"
}

// https://gobyexample.com/channel-buffering
func main() {
	messages := make(chan string, 2)
	go addMessage(messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
