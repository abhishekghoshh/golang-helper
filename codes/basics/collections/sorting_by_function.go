package collections

import (
	"fmt"
	"sort"
)

// https://gobyexample.com/sorting-by-functions
func SortingByFunction() {
	fruits := StringWrapper([]string{"peach", "banana", "kiwi"})
	sort.Sort(fruits)
	fmt.Println(fruits)
}

type StringWrapper []string

func (s StringWrapper) Len() int {
	return len(s)
}
func (s StringWrapper) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s StringWrapper) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
