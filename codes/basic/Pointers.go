package main

import "fmt"

func changeToNum(num, val int) {
	num = val
}
func changeToNumWithPointer(ptr *int, val int) {
	*ptr = val
}
func main() {
	x := 12
	p := &x
	*p = 10
	fmt.Println(p, " -> ", x)

	changeToNum(x, 8)
	fmt.Println(p, " -> ", x)

	changeToNumWithPointer(p, 8)
	fmt.Println(p, " -> ", x)
}
