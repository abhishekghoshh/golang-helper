package introduction

import "fmt"

/*
https://en.wikipedia.org/wiki/Recursion_(computer_science)
https://gobyexample.com/recursion

This fact function calls itself until it reaches the base case of fact(0).
*/

func Recursion() {
	fmt.Println("factorial of 7 is", fact(7))

	var fib func(n int) int

	// inline function
	// Closures can also be recursive, but this requires the closure
	// to be declared with a typed var explicitly before it’s defined.
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	// Since fib was previously declared in main, Go knows which function to call with fib here.
	fmt.Println("7th fibonacci number is", fib(7))
}

func fact(n int) int {
	if n == 1 {
		return n
	}
	return n * fact(n-1)
}
