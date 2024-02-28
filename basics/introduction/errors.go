package introduction

import (
	"errors"
	"fmt"
)

/*
https://gobyexample.com/errors
https://go.dev/blog/error-handling-and-go


In Go it's idiomatic to communicate errors via an explicit, separate return value.
This contrasts with the exceptions used in languages like Java and Ruby
and the overloaded single result / error value sometimes used in C. Go's approach makes it easy to see which functions
return errors and to handle them using the same language constructs employed for any other, non-error tasks.

By convention, errors are the last return value and have type error, a built-in interface.
	type error interface {
		Error() string
	}
errors.New constructs a basic error value with the given error message.
A nil value in the error position indicates that there was no error.

unctions often return an error value, and calling code should handle errors by testing whether the error equals nil.
	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)
*/

func Errors() {
	for _, i := range []int{7, 42} {
		// Note that the use of an inline error check on the if line is a common idiom in Go code.
		if r, e := isNotFourtyTwo(i); e != nil {
			fmt.Println("isFourtyTwo failed:", e)
		} else {
			fmt.Println("isFourtyTwo worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := isNotFourtyTwoElseThrowCustomError(i); e != nil {
			fmt.Println("isNotFourtyTwoElseThrowCustomError failed:", e)
		} else {
			fmt.Println("isNotFourtyTwoElseThrowCustomError worked:", r)
		}
	}

	// If you want to programmatically use the data in a custom error,
	// you'll need to get the error as an instance of the custom error type via type assertion.
	_, e := isNotFourtyTwoElseThrowCustomError(42)
	if ae, ok := e.(*customArgError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

func isNotFourtyTwo(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type customArgError struct {
	arg  int
	prob string
}

func (e *customArgError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func isNotFourtyTwoElseThrowCustomError(arg int) (int, error) {
	if arg == 42 {
		// In this case we use &customArgError syntax to build a new struct,
		// supplying values for the two fields arg and message.
		return -1, &customArgError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
