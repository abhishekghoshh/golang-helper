package concurrency

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

/*
https://gobyexample.com/channel-buffering
https://www.youtube.com/watch?v=tSjnf6l8cq8


By default channels are unbuffered, meaning that they will only accept sends (chan <-)
if there is a corresponding receive (<- chan) ready to receive the sent value.
Buffered channels accept a limited number of values without a corresponding receiver for those values.
*/

func ChannelBuffering() {
	go func() {
		for {
			printStats()
		}
	}()

	// Here we make a channel of strings buffering up to 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// just run the below example and see the output of the number of goroutines
	// We can also use buffer channel to limit the number of goroutines that can run at a time.
	eventManager := New(100)
	// with this event limiter we are blocking the the creation of the new goroutines inside the for loop
	// it will not create more
	eventLimiter := make(chan interface{}, 5)
	for event := range eventManager.Stream() {
		eventLimiter <- true // we can anything
		go func() {
			eventManager.process(event)
			<-eventLimiter
		}()
	}

}

type EventManager struct {
	ch    chan int
	count int
}

func New(c int) *EventManager {
	ch := make(chan int, 10)
	return &EventManager{ch, c}
}

func (em *EventManager) Stream() chan int {
	c := em.count
	go func() {
		for i := 1; i <= c; i++ {
			time.Sleep(100 * time.Millisecond)
			em.ch <- i
		}
		close(em.ch)
	}()
	return em.ch
}

func (em *EventManager) process(id int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Processing event", id, "on goroutine", em.goid())
}

func (em *EventManager) goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// The first line of stack trace contains: "goroutine <id> [running]:"
	fields := bytes.Fields(buf[:n])
	id, err := strconv.Atoi(string(fields[1]))
	if err != nil {
		return -1 // Return -1 if there's an error
	}
	return id
}

func printStats() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Println("========== Runtime Stats ==========")
	fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("Heap Alloc: %d KB\n", memStats.HeapAlloc/1024)
	fmt.Printf("Total Alloc: %d KB\n", memStats.TotalAlloc/1024)
	fmt.Printf("Heap Objects: %d\n", memStats.HeapObjects)
	fmt.Printf("GC Cycles: %d\n", memStats.NumGC)
	fmt.Println("===================================")
	time.Sleep(time.Second)
}
