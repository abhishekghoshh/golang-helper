package introduction

import "fmt"

func Generics() {
	printAll(1, 2, 3, 4)
	printAll("Abhishek", "Nasim", "Bishal")

	DoWork(worker("Abhishek"))

	fmt.Println(Equal(1, 4))

	var m = map[int]string{
		1: "2",
		2: "4",
		4: "8",
	}
	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)
}

func printAll[T any](args ...T) {
	for _, v := range args {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

type Worker interface {
	Work()
}

type worker string

func (w worker) Work() {
	fmt.Printf("%s is working\n", w)
}

func DoWork[T Worker](things ...T) {
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
