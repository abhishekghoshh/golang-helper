package main

import (
	"errors"
	"fmt"
)

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
		return -1, &customArgError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

// https://gobyexample.com/errors
func main() {
	for _, i := range []int{7, 42} {
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

	_, e := isNotFourtyTwoElseThrowCustomError(42)
	if ae, ok := e.(*customArgError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
