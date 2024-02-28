package collections

import "fmt"

/*
range iterates over elements in a variety of data structures.
Let's see how to use range with some of the data structures we've already learned.

range on arrays and slices provides both the index and value for each entry.
We didn't need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though.
*/
func Ranges() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for index, num := range nums {
		fmt.Printf("nums[%d] = %d \n", index, num)
	}

	arr := make([]int, 3)

	for i, item := range arr {
		fmt.Println(i, "->[zero value]->", item)
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and the second the rune itself
	x := "hello"
	for _, c := range x {
		fmt.Println(c)
	}

	for _, c := range x {
		fmt.Printf("%c ", c)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}
}
