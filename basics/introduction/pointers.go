package introduction

import "fmt"

/*
https://en.wikipedia.org/wiki/Pointer_(computer_programming)
https://gobyexample.com/pointers

Go supports pointers, allowing you to pass references to values and records within your program.
how pointers work in contrast to values with 2 functions: zeroval and zeroptr. zeroval has an int parameter,
so arguments will be passed to it by value. zeroval will get a copy of ival distinct from the one in the calling function.

	func zeroval(ival int) {
		ival = 0
	}
	func zeroptr(iptr *int) {
		*iptr = 0
	}

zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
The *iptr code in the function body then dereferences the pointer from its memory address to
the current value at that address. Assigning a value to a dereferenced pointer changes the value at the referenced address.

The &i syntax gives the memory address of i, i.e. a pointer to i.

zeroval doesn't change the i in main, but zeroptr does because it has a reference to the memory address for that variable.
*/

func Pointers() {
	x := 12
	p := &x
	*p = 10
	fmt.Println(p, " -> ", x)

	changeToNum(x, 8)
	fmt.Println(p, " -> ", x)

	changeToNumWithPointer(p, 8)
	fmt.Println(p, " -> ", x)

	var intPtr *int
	fmt.Println(intPtr)
}

func changeToNum(num, val int) {
	fmt.Println("Value of num before assignment is", num)
	num = val
	fmt.Println("Value of num after assignment is", num)
}
func changeToNumWithPointer(ptr *int, val int) {
	*ptr = val
}
