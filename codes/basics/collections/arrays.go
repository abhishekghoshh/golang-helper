package collections

import "fmt"

/*
 Here we create an array a that will hold exactly 5 ints. The type of elements and length are both part of the array's type.
 By default an array is zero-valued, which for ints means 0s.
 We can set a value at an index using the array[index] = value syntax, and get a value with array[index].
 The builtin len returns the length of an array.

 Use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}

 Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
*/

func Arrays() {
	// Now a is an array of 5 integers.
	var arr [5]int
	arr[0] = 55
	fmt.Println(arr)

	arr1 := [5]int{0, 2, 4, 6, 8}
	fmt.Println(arr1[1])

	fmt.Println(arr1[2:4])

	// this is a two dimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// this is also a slice
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

	// array of the custom types
	type Person struct {
		name string
		age  int
	}
	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}
	fmt.Println(people)

	// we can also remove the Person tag on each element, it is inferred
	// that array of personal will consits the person type, so we can remove that

	people = []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}
	fmt.Println(people)

}
