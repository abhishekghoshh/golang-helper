package introduction

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

/*
Go has only one looping construct, the for loop.

The init and post statements are optional.

	for ; sum < 1000; {
		sum += sum
	}

At that point you can drop the semicolons: C's while is spelled for in Go.

	for sum < 1000 {
		sum += sum
	}

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

Like for, the if statement can start with a short statement to execute before the condition.

	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim

Variables declared inside an if short statement are also available inside any of the else blocks.
Variables declared by the statement are only in scope until the end of the if.
*/
func FlowControl() {

	fmt.Println("Print 1 to 10")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// without init and post statement
	// this is working as the shile loop
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("sum is", sum)

	// This is an infinte loop
	num := 0
	for {
		if num == 10 {
			break
		}
		num++
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// Sample switch case
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Another switch example
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// Switch without a condition is the same as switch true. This construct can be a clean way to write long if-then-else chains.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
