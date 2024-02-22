package collections

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

/*
https://gobyexample.com/sorting-by-functions

Sometimes we'll want to sort a collection by something other than its natural order.
For example, suppose we wanted to sort strings by their length instead of alphabetically. Here's an example of custom sorts in Go.
*/
func CustomSorting() {

	fruits := Fruits{
		"peach",
		"banana",
		"kiwi",
	}
	sort.Sort(fruits)
	fmt.Println(fruits)

	fruits = Fruits{
		"peach",
		"banana",
		"kiwi",
	}
	// slices sorted the fruits on alphabetical order
	slices.Sort(fruits)
	fmt.Println(fruits)

	// custom comparator function
	// We implement a comparison function for string lengths. cmp.Compare is helpful for this.
	comparator := func(fruit1, fruit2 string) int {
		return cmp.Compare(len(fruit1), len(fruit2))
	}
	// Now we can call slices.SortFunc with this custom comparison function to sort fruits by name length.
	slices.SortFunc(fruits, comparator)
	fmt.Println(fruits)

	// We can use the same technique to sort a slice of values that aren't built-in types.
	type Person struct {
		name string
		age  int
	}
	people := []Person{
		{name: "Abhishek", age: 37},
		{name: "Nasim", age: 25},
		{name: "Sayan", age: 72},
	}

	slices.SortFunc(people,
		func(person1, person2 Person) int {
			return cmp.Compare(person1.age, person2.age)
		},
	)
	fmt.Println(people)
}

type Fruits []string

func (fruits Fruits) Len() int {
	return len(fruits)
}
func (fruits Fruits) Swap(i, j int) {
	fruits[i], fruits[j] = fruits[j], fruits[i]
}
func (fruits Fruits) Less(i, j int) bool {
	return len(fruits[i]) < len(fruits[j])
}
