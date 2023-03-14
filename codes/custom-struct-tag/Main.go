package main

import (
	"student/student"
)

func main() {
	validate := &student.Student{Age: 20, Name: "Abhishek Ghosh"}
	validate.Validate()
}
