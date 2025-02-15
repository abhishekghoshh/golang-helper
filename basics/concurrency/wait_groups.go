package concurrency

import (
	"fmt"
	"sync"
	"time"
)

/*
https://gobyexample.com/waitgroups
https://pkg.go.dev/golang.org/x/sync/errgroup

To wait for multiple goroutines to finish, we can use a wait group.
*/
func WaitGroups() {

	// This is the function we'll run in every goroutine.
	var worker = func(id int) {
		fmt.Printf("Worker %d starting\n", id)

		// Sleep to simulate an expensive task.
		time.Sleep(time.Second)
		fmt.Printf("Worker %d done\n", id)
	}

	// This WaitGroup is used to wait for all the goroutines launched here to finish.
	// Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// Launch several goroutines and increment the WaitGroup counter for each.
		wg.Add(1)

		// Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
		// This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.
		go func(num int) {
			defer wg.Done()
			worker(num)
		}(i)
	}

	// Block until the WaitGroup counter goes back to 0; all the workers notified they're done.
	wg.Wait()

	// Note that this approach has no straightforward way to propagate errors from workers.
	// For more advanced use cases, consider using the errgroup package.

	fmt.Print("Another producer consumer example\n")
	////////////////////////////////////////
	// We will mimic here a producer consumer pattern using wait groups
	var wg1 sync.WaitGroup

	readerCount := 5
	channelBufferSize := 1
	ch := make(chan int, channelBufferSize)

	consumer := func(id int) {
		defer wg1.Done()
		for i := range ch {
			time.Sleep(time.Millisecond)
			fmt.Printf("Reader %d: %d\n", id, i)
		}
	}
	// Start the consumers
	go func() {
		for i := 1; i <= readerCount; i++ {
			wg1.Add(1)
			go consumer(i)
		}
	}()

	producer := func(count int) {
		for i := 0; i < count; i++ {
			time.Sleep(time.Millisecond)
			ch <- i
		}
		close(ch)
	}
	producer(100)
	wg1.Wait()
}
