package introduction

import "fmt"

/*
clousure is a function returning another function
while the calling funtion will maintain its state
*/

// https://gobyexample.com/closures
func Clousures() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
