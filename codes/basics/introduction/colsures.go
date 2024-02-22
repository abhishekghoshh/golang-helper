package introduction

import "fmt"

/*
https://en.wikipedia.org/wiki/Anonymous_function
https://en.wikipedia.org/wiki/Closure_(computer_programming)
// https://gobyexample.com/closures
clousure is a function returning another function, while the calling funtion will maintain its state

Go supports anonymous functions, which can form closures.
Anonymous functions are useful when you want to define a function inline without having to name it.


*/

func Clousures() {

	// This function intSeq returns another function, which we define anonymously in the body of intSeq.
	// The returned function closes over the variable i to form a closure.
	// We call intSeq, assigning the result (a function) to nextInt.
	// This function value captures its own i value, which will be updated each time we call nextInt.
	nextIntFn1 := intSeq()

	fmt.Println(nextIntFn1())
	fmt.Println(nextIntFn1())
	fmt.Println(nextIntFn1())

	// To confirm that the state is unique to that particular function, create and test a new one.
	nextIntFn2 := intSeq()
	fmt.Println(nextIntFn2())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
