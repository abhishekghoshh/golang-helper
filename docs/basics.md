# Go Basics

## Sample Go Code

```go
package main // Package declaration

import "fmt" // Import other packages that we need

func main() {  //  Declare functions, tell Go to do things
    fmt.Println("Hello World")
}
```

To import a package, we use the `import` statement. One of the most popular packages is `"fmt"`, which stands for format, and provides input and output functionality.
Package `main` is the entry point for go program.
You can import multiple packages at once, using parentheses. For example: `import ("fmt" "math")`.

In Go, the var keyword is used to declare variables.
For example: `var i int`

```go
var i int = 8
fmt.Println(i) 
```

## Variables

The code above declares a variable named i of type int. int stands for integer and represents whole numbers.
We can assign the variable a value and output it.
You can also declare multiple variables on one line and assign values to them:

```go
var i, j int = 8, 42
fmt.Println(i) 
fmt.Println(j)
```

If you assign a value to the variable, you can omit the type declaration, as Go will automatically take the type of the value assigned.
Go supports short variable declaration using `:=`:
For example: `i := 8`

The variable `i` will be an integer variable and the value of `i` will be `8`.
It can also be used to declare and initialize multiple variables on one line: `x, y := 10, 5`.

## Data Types

We used int to declare integer values in our previous lesson. `var i int = 8`

| Type | Description |
|---|---|
| `bool` | true or false |
| `string` | immutable sequence of bytes |
| `int` | 32 or 64 bits depending on system |
| `uint` | unsigned integer |
| `int8..int64` | explicitly sized signed integers |
| `uint8..uint64` | explicitly sized unsigned integers |
| `byte` | alias for `uint8` |
| `rune` | alias for `int32`, represents a Unicode code point |
| `float32`, `float64` | floating point |
| `complex64`, `complex128` | complex numbers |

```go
var x int = 42
var y float32 = 1.37
var name string = "James"
var online bool = true
```

Another interesting feature of Go are **zero values**: variables that are declared without a value take the zero value of their type. `0` for numeric types, `false` for the boolean type, `""` for strings.

```go
var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
)

fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)

var i int     // zero value: 0
var f float64 // zero value: 0
var b bool    // zero value: false
var s string  // zero value: ""
fmt.Printf("zero values: %v %v %v %q\n", i, f, b, s)
```

## Constants

A variable can change its value during the program but in some cases your program may need values that are preserved during the program. These are called constants and they cannot be changed from their initial value.

Constants are declared like variables, but with the `const` keyword and need to be assigned a value:

```go
const Pi = 3.14 
fmt.Println(Pi)
```

Now, `Pi` is a constant and cannot be changed. Constants cannot be declared using the `:=` syntax.

## Input

The `"fmt"` package also allows you to take input from the user of the program.

```go
var input int = 10
fmt.Scanln(&input)
fmt.Println("input is", input)
```

*Note the ampersand & before the variable name -- it is used to return the address of a variable.*
