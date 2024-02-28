package introduction

import "fmt"

/*
https://gobyexample.com/panic
https://gobyexample.com/recover

A panic typically means something went unexpectedly wrong.
Mostly we use it to fail fast on errors that shouldn't occur during normal operation,
or that we aren't prepared to handle gracefully

A common use of panic is to abort if a function returns an error value that we don't know how to (or want to) handle.
Running this program will cause it to panic, print an error message and goroutine traces, and exit with a non-zero status.
When first panic in main fires, the program exits without reaching the rest of the code.
Note that unlike some languages which use exceptions for handling of many errors,
in Go it is idiomatic to use error-indicating return values wherever possible.

Go makes it possible to recover from a panic, by using the recover built-in function.
A recover can stop a panic from aborting the program and let it continue with execution instead.
An example of where this can be useful: a server wouldn't want to crash if one of the client connections
exhibits a critical error. Instead, the server would want to close that connection and continue serving other clients.
In fact, this is what Go's net/http does by default for HTTP servers.
*/

// This code will not run, because mayPanic panics.
// The execution of main stops at the point of the panic and resumes in the deferred closure.
func PanicAndRecover() {

	// recover must be called within a deferred function. When the enclosing function panics,
	// the defer will activate and a recover call within it will catch the panic
	defer func() {
		// The return value of recover is the error raised in the call to panic.
		if r := recover(); r != nil {
			fmt.Println("Recovered from a panic. \nPanic reasone is:", r)
		}
	}()
	Panic()
	fmt.Println("After Panic()")
}

func Panic() {
	panic("This is a custom panic")
}
