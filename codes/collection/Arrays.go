package main

import "fmt"

func main() {
	//Now a is an array of 5 integers.
	var arr [5]int
	arr[0] = 55
	fmt.Println(arr)

	//this is a slice
	arr1 := [5]int{0, 2, 4, 6, 8}
	fmt.Println(arr1[1])

	fmt.Println(arr1[2:4])

	arr2 := make([]int, 5)
	fmt.Println(arr2)
	arr2 = append(arr2, 8)
	fmt.Println(arr2)

	// this is a slice
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
