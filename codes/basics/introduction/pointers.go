package introduction

import "fmt"

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
	fmt.Println("Value of num is", num)
	num = val
	fmt.Println("Value of num is", num)
}
func changeToNumWithPointer(ptr *int, val int) {
	*ptr = val
}
