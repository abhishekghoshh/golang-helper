package main

import (
	"fmt"
	"reflect"
)

func Reflections() {
	x := 100
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Println("Type:", t) // "Type: int"
}
