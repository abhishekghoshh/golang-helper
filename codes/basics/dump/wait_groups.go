package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// https://gobyexample.com/waitgroups
func WaitGroups() {

	var worker = func(id int) {
		fmt.Printf("Worker %d starting\n", id)

		time.Sleep(time.Second)
		fmt.Printf("Worker %d done\n", id)
	}

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			worker(num)
		}(i)
	}
	wg.Wait()

}
