package collections

import "fmt"

/*
Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
An uninitialized slice equals to nil and has length 0.

To create an empty slice with non-zero length, use the builtin make.
Here we make a slice of strings of length 3 (initially zero-valued).
By default a new slice's capacity is equal to its length; if we know the slice is going to grow ahead of time,
it's possible to pass a capacity explicitly as an additional parameter to make.

len returns the length of the slice as expected.

In addition to these basic operations, slices support several more that make them richer than arrays.
One is the builtin append, which returns a slice containing one or more new values. Note that we need
to accept a return value from append as we may get a new slice value.


Blog post : https://go.dev/blog/slices-intro
*/

func Slices() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be copy"d. Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a "slice" operator with the syntax slice[low:high].
	// For example, this gets a slice of the elements s[2], s[3], and s[4].
	// This slices up to (but excluding) s[5].
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// We can declare and initialize a variable for slice in a single line as well
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
