package concurrency

import "fmt"

// https://gobyexample.com/channel-directions
func ChannelDirections() {
	var ping = func(pings chan<- string, names []string) {
		for _, name := range names {
			pings <- name
		}
		close(pings)
	}

	var pong = func(pings <-chan string, pongs chan<- string) {
		for name := range pings {
			pongs <- name
		}
		close(pongs)
	}

	names := []string{"Abhishek Ghosh", "Nasim Molla", "Bishal Mukherjee", "Abhishek Pal", "Avishek Saha", "Niket Saha"}
	pings := make(chan string)
	pongs := make(chan string)
	go ping(pings, names)
	go pong(pings, pongs)
	for name := range pongs {
		fmt.Println(name)
	}
}
