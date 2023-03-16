package main

import (
	"fmt"
	"student/student"
)

func main() {
	validate := &student.Student{Age: 25, Name: ""}
	err := validate.Validate()
	fmt.Println(err)
}
