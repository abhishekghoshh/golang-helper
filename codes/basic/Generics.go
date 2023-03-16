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

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func main() {
	print(1, 2, 3, 4)
	print("Abhishek", "Nasim", "Bishal")

	DoWork(worker("Abhishek"))

	fmt.Println(Equal(1, 4))

	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys:", MapKeys(m))
	_ = MapKeys[int, string](m)
}
