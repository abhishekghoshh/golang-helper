package concurrency

import "fmt"

// https://gobyexample.com/closing-channels
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func(ch chan int) {
		for {
			j, more := <-ch
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}(jobs)

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
