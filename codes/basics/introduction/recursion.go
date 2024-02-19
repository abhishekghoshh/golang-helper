package introduction

import "fmt"

func Recursion() {
	fmt.Println("factorial of 7 is", fact(7))

	var fib func(n int) int

	// inline function
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println("7th fibonacci number is", fib(7))
}

func fact(n int) int {
	if n == 1 {
		return n
	}
	return n * fact(n-1)
}
