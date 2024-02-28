package collections

import (
	"fmt"
	"slices"
	"sort"
)

/*
https://gobyexample.com/sorting
https://pkg.go.dev/cmp#Ordered

Go's slices package implements sorting for builtins and user-defined types. We'll look at sorting for builtins first.

Sorting functions are generic, and work for any ordered built-in type. For a list of ordered types, see cmp.Ordered.
We can also use the slices package to check if a slice is already in sorted order.
*/
func Sorting() {
	alphablet := []string{"c", "a", "b"}
	// sorting on the alphabet string, using the sort package on Strings method
	sort.Strings(alphablet)
	fmt.Println("alphablets : ", alphablet)

	alphablet = []string{"c", "a", "b"}
	// sorting on the alphabet string, using the slices package sort method
	slices.Sort(alphablet)
	fmt.Println("alphablet : ", alphablet)

	// using Ints method from sorted package
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints : ", ints)

	// using general sort method for the generic list
	ints = []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// is sorted specifically for the int
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	// general is sorted method for a generic list
	s = slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
