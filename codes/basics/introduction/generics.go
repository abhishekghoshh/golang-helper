package introduction

import "fmt"

/*
Go functions can be written to work on multiple types using type parameters.
The type parameters of a function appear between brackets, before the function's arguments.
	func Index[T comparable](s []T, x T) int

This declaration means that s is a slice of any type T that fulfills the built-in constraint comparable.
x is also a value of the same type.
comparable is a useful constraint that makes it possible to use the == and != operators on values of the type.
In this example, we use it to compare a value to all slice elements until a match is found.
This Index function works for any type that supports comparison.


In addition to generic functions, Go also supports generic types.
A type can be parameterized with a type parameter,
which could be useful for implementing generic data structures.
This example demonstrates a simple type declaration for a singly-linked list holding any type of value.
As an exercise, add some functionality to this list implementation.

LinkedList represents a singly-linked list that holds values of any type.
	type LinkedList[T any] struct {
		next *LinkedList[T]
		val  T
	}

*/

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

	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
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

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}
