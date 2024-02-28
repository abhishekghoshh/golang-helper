package concurrency

import (
	"fmt"
	"time"
)

/*
https://gobyexample.com/timeouts

Timeouts are important for programs that connect to external resources or that otherwise
need to bound execution time. Implementing timeouts in Go is easy and elegant thanks to channels and select.

For our example, suppose we're executing an external call that returns its result
on a channel c1 after 2s. Note that the channel is buffered, so the send in the goroutine is nonblocking.
This is a common pattern to prevent goroutine leaks in case the channel is never read.

Here's the select implementing a timeout. res := <-c1 awaits the result and <-time.After
awaits a value to be sent after the timeout of 1s. Since select proceeds with the first receive
that's ready, we'll take the timeout case if the operation takes more than the allowed 1s.
*/
func Timeouts() {

	c1 := make(chan string, 1)
	go httpGetCallToApi(c1, "result 1")

	// timeout is set to 1 sec, so we are hitting the timeout here
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go httpGetCallToApi(c2, "result 2")

	// the ttimeout is set to 3 sec here, so we are not hitting the timeout,
	// we are getting the value from the channel
	// If we allow a longer timeout of 3s, then the receive from c2 will succeed and we'll print the result.
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

	// Running this program shows the first operation timing out and the second succeeding.
}

func httpGetCallToApi(ch chan string, message string) {
	time.Sleep(2 * time.Second)
	ch <- message
}
