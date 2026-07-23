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

## Complete Examples

```go
func main() {
    // Array of 5 ints (fixed size)
    var arr [5]int
    arr[0] = 55

    arr1 := [5]int{0, 2, 4, 6, 8}
    fmt.Println(arr1[1])
    fmt.Println(arr1[2:4]) // slice from array

    // 2D array
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }

    // Slice: dynamically-sized
    s := make([]string, 3)
    s[0] = "a"; s[1] = "b"; s[2] = "c"
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("len:", len(s))

    // Copy slice
    c := make([]string, len(s))
    copy(c, s)

    // Slice operator: slice[low:high]
    l := s[2:5]  // elements 2,3,4
    l = s[:5]    // first 5 elements
    l = s[2:]    // from index 2 to end

    // Inline slice declaration
    t := []string{"g", "h", "i"}

    // 2D slice (inner slices can vary)
    twoDSlice := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoDSlice[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoDSlice[i][j] = i + j
        }
    }
}
```

## Sorting

```go
// sort package
strs := []string{"c", "a", "b"}
sort.Strings(strs)

ints := []int{7, 2, 4}
sort.Ints(ints)

// slices package (generic)
slices.Sort(strs)
slices.Sort(ints)

// check if sorted
sort.IntsAreSorted(ints)
slices.IsSorted(ints)

// custom sorting with comparator
cmpFn := func(a, b string) int {
    return cmp.Compare(len(a), len(b))
}
slices.SortFunc(fruits, cmpFn)

// custom sorting by implementing sort.Interface
type Fruits []string
func (f Fruits) Len() int            { return len(f) }
func (f Fruits) Swap(i, j int)       { f[i], f[j] = f[j], f[i] }
func (f Fruits) Less(i, j int) bool  { return len(f[i]) < len(f[j]) }
sort.Sort(fruits)
```

## Linked List (Generic)

```go
type List[T any] struct {
    head, tail *Node[T]
}

type Node[T any] struct {
    val  T
    next *Node[T]
}

func NewList[T any]() List[T] { return List[T]{} }

func (lst *List[T]) Push(v T) {
    node := &Node[T]{val: v}
    if lst.tail == nil {
        lst.head = node
    } else {
        lst.tail.next = node
    }
    lst.tail = node
}

func (lst *List[T]) GetAll() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}
```

## Ordered Map (Generic)

Preserves insertion order using a key slice alongside a map.

```go
type OrderedMap[K comparable, V any] struct {
    store map[K]V
    keys  []K
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
    return &OrderedMap[K, V]{store: map[K]V{}, keys: []K{}}
}

func (om *OrderedMap[K, V]) Set(key K, val V) {
    if _, exists := om.store[key]; !exists {
        om.keys = append(om.keys, key)
    }
    om.store[key] = val
}

func (om *OrderedMap[K, V]) Delete(key K) {
    delete(om.store, key)
    for i, k := range om.keys {
        if k == key {
            om.keys = append(om.keys[:i], om.keys[i+1:]...)
            break
        }
    }
}

func (om *OrderedMap[K, V]) Iterator() func() (*int, *K, V) {
    index := 0
    return func() (_ *int, _ *K, _ V) {
        if index > len(om.keys)-1 { return }
        row := om.keys[index]
        index++
        return &[]int{index-1}[0], &row, om.store[row]
    }
}
```

## Queue

```go
type Queue struct { items []int }

func (q *Queue) Enqueue(item int) {
    q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
    if len(q.items) < 1 { return -1 }
    item := q.items[0]
    q.items = q.items[1:]
    return item
}
```
