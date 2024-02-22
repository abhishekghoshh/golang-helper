package concurrency

import (
	"fmt"
	"time"
)

// TODO check it later
// https://gobyexample.com/timers
func Timers() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func(timer *time.Timer) {
		<-timer.C
		fmt.Println("Timer 2 fired")
	}(timer2)

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
