package concurrency

import "fmt"

/*
https://gobyexample.com/channel-directions
*/

func ChannelDirections() {

	// This ping function only accepts a channel for sending values.
	// It would be a compile-time error to try to receive on this channel.
	var ping = func(pings chan<- string, names []string) {
		for _, name := range names {
			pings <- "Hi ping," + name
		}
		close(pings)
	}

	// The pong function accepts one channel for receives (pings) and a second for sends (pongs).
	var pong = func(pings <-chan string, pongs chan<- string) {
		for name := range pings {
			pongs <- name + ", Hi pong"
		}
		close(pongs)
	}

	names := []string{
		"Abhishek Ghosh",
		"Nasim Molla",
		"Bishal Mukherjee",
		"Abhishek Pal",
		"Avishek Saha",
		"Niket Saha",
	}
	pings := make(chan string)
	pongs := make(chan string)
	go ping(pings, names)
	go pong(pings, pongs)

	// we are waiting for the pong channel for sending the data
	for incomingName := range pongs {
		fmt.Println(incomingName)
	}
}
