package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// https://gobyexample.com/atomic-counters
func AtomicCounters() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(num *uint64) {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(num, 1)
			}
			wg.Done()
		}(&ops)
	}

	wg.Wait()
	fmt.Println("ops:", ops)
}
