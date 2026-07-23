# Arrays, Slices, and Maps

## Arrays

An array is a sequence of elements of the same type. An array is defined using square brackets which define the number of elements the array will hold.

An array type definition specifies a length and an element type. For example, the type `[4]int` represents an array of four integers. An array's size is fixed; its length is part of its type (`[4]int` and `[5]int` are distinct, incompatible types).

Arrays can be indexed in the usual way, so the expression `s[n]` accesses the nth element, starting from zero.
Arrays do not need to be initialized explicitly; the zero value of an array is a ready-to-use array whose elements are themselves zeroed.

The in-memory representation of `[4]int` is just four integer values laid out sequentially.

**Go's arrays are values.** An array variable denotes the entire array; it is not a pointer to the first array element (as would be the case in C). This means that when you assign or pass around an array value you will make a copy of its contents. (To avoid the copy you could pass a pointer to the array, but then that's a pointer to an array, not an array.) One way to think about arrays is as a sort of struct but with indexed rather than named fields: a fixed-size composite value.

For example:

```go
var a [5]int
```

Now `a` is an array of 5 integers.

We can also define and initialize values of the array using the following syntax:

```go
a := [5]int{0, 2, 4, 6, 8}
```

As we can see, we need to provide the size of the array when declaring it. This means that arrays have a fixed size.
After declaring the array, we can access its elements using square brackets and their indexes.
For example:

```go
var a [5]int
a[0] = 8
a[1] = 42
fmt.Println(a[1])
```

Each element of an array has its unique index, starting with 0. The first element has the index 0, the second element has the index 1, and so on. We can assign values to array elements, as well as retrieve their value. By default, the elements of the array are initialized to the zero value of the given type.

## Slices

The `slice` type is an abstraction built on top of Go's array type.

An array has a fixed size, meaning once defined, you cannot change the number of elements it holds. To overcome this, Go provides the `slice`, which is a **dynamically-sized view into the elements of an array**.

A slice is based on an array and is defined by specifying two indices, a low and high bound, separated by a colon:

```go
a := []int{0, 2, 4, 6, 8}
```

A slice does not store any data; it just describes a section of an underlying array. Changing the elements of a slice modifies the corresponding elements of its underlying array.

Arrays have their place, but they're a bit inflexible, so you don't see them too often in Go code. Slices, though, are everywhere. They build on arrays to provide great power and convenience.

The type specification for a slice is `[]T`, where `T` is the type of the elements of the slice. Unlike an array type, a slice type has no specified length.

**A slice literal is declared just like an array literal, except you leave out the element count.**

We can have multiple slices of the same array. A change in any of them will be seen in all slices, as it will affect the underlying array.

### Creating Slices with make

A slice can be created with the built-in function called `make`, which has the signature:

```go
func make([]T, len, cap) []T
a := make([]int, 5) 
```

The length and capacity of a slice can be inspected using the built-in `len` and `cap` functions.

The make function creates an array of the given type and size, and returns a slice that refers to that array. After creating a slice, we can add/append new elements to it using the `append()` function:

```go
a := make([]int, 5)
a = append(a, 8)
fmt.Println(a)
```

The `append()` function takes the slice as its first argument and the elements to be added to the end of the slice as its next argument. It then returns a new slice, containing the old slice plus the new elements appended. We can append multiple values at once by just comma-separating the values as arguments, for example:
`append(s, 1, 2, 3)`

The zero value of a slice is nil. The `len` and `cap` functions will both return 0 for a nil slice.

A slice can also be formed by **"slicing"** an existing slice or array. Slicing is done by specifying a half-open range with two indices separated by a colon. For example, the expression `b[1:4]` creates a slice including elements 1 through 3 of b (the indices of the resulting slice will be 0 through 2).

For more reference check: [Go Slices Intro](https://go.dev/blog/slices-intro)

## Range

Now that we know how to create arrays and slices, let's learn how to iterate over their elements using a loop. The range form of the for loop allows you to iterate over a slice. During each iteration of the loop, it returns two values: the index of the element and its value.

```go
func main() {
    a := make([]int, 5)
    a[1] = 2
    a[2] = 3
    for i, v := range a {
        fmt.Println(i, v)
    }
}
```

If you want only the values, you can skip the index using an underscore `_`:

```go
func main() {
    a := make([]int, 5)
    a[1] = 2
    a[2] = 3
    for _, v := range a {
        fmt.Println(v)
    }
}
```

You can use ranges for slices as well as arrays. The range can also be used to iterate over the characters of a string:

```go
func main() {
    x := "hello"
    for _, c := range x {
        fmt.Print(c)
    }
}
```

```go
func main() {
    x := "hello"
    for _, c := range x {
        fmt.Printf("%c ", c)
    }
}
```

The `Printf()` function is similar to the one in C, taking the format of the output as its argument. `%c` in our case denotes a character, while `\n` defines a new line.

## Maps

Maps are used to store `key:value` pairs. The key is always unique. We can create a map using the `make()` function, similar to arrays.

```go
freq := make(map[string]int)
```

You can also initialize a map using the following syntax. **If the requested key does not exist in the map, a zero value will be returned**. Maps are also called dictionaries, associative arrays, or hash tables. You can use the `delete` function to remove an element from the map. Maps are printed in the form `map[key:value key:value]` when output with `fmt.Println()`.

```go
func main() {
    m := map[string]int{
        "James": 42,
        "Amy":   24}
    fmt.Println(m["Amy"])
}
```

```go
func main() {
    m := map[string]int{
        "James": 42,
        "Amy":   24}
    delete(m, "James")
    fmt.Println(m)
}
```
