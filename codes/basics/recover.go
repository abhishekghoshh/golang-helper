package main

import "fmt"

func Panic() {
	panic("a problem")
}

// https://gobyexample.com/panic
// https://gobyexample.com/recover
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:", r)
		}
	}()
	Panic()
	fmt.Println("After Panic()")
}
