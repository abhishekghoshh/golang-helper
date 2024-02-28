package utility

import (
	"fmt"
	"reflect"
)

/*
https://go101.org/article/reflection.html
*/

func Reflections() {
	x := 100
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Println("Type:", t) // "Type: int"

	// there is many more
}
