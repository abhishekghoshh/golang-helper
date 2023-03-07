package main

import "fmt"

func print[T any](args ...T) {
	for _, v := range args {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

type Person interface {
	Work()
}

type worker string

func (w worker) Work() {
	fmt.Printf("%s is working\n", w)
}

func DoWork[T Person](things ...T) {
	for _, v := range things {
		v.Work()
	}
}
func Equal[T comparable](a, b T) bool {
	return a == b
}

func main() {
	print(1, 2, 3, 4)
	print("Abhishek", "Nasim", "Bishal")

	DoWork(worker("Abhishek"))

	fmt.Println(Equal(1, 4))
}
