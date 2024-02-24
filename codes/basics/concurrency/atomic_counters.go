package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
https://gobyexample.com/atomic-counters

The primary mechanism for managing state in Go is communication over channels.
We saw this for example with worker pools. There are a few other options for managing state though.
Here we'll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
*/

func AtomicCounters() {

	// We'll use an atomic integer type to represent our (always-positive) counter.
	var ops atomic.Uint64

	var normal_int int = 0

	// A WaitGroup will help us wait for all goroutines to finish their work.
	var wg sync.WaitGroup

	// We'll start 50 goroutines that each increment the counter exactly 1000 times.
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
				normal_int++
			}
			wg.Done()
		}()
	}

	// Wait until all the goroutines are done.
	wg.Wait()

	// Here no goroutines are writing to 'ops', but using Load it's safe to atomically
	// read a value even while other goroutines are (atomically) updating it.
	fmt.Println("ops:", ops.Load())

	// value of normal int will never be 50000 because it is not thread safe
	fmt.Println("value of normal int:", normal_int)

	// We expect to get exactly 50,000 operations. Had we used a non-atomic integer and incremented it with ops++,
	// we'd likely get a different number, changing between runs, because the goroutines would interfere with each other.
	// Moreover, we'd get data race failures when running with the -race flag.
}
