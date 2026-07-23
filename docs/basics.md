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
Let's see what other common types Go supports.

- **float32** - a single-precision floating point value.
- **float64** - a double-precision floating point value.
- **string** - a text value.
- **bool** - Boolean true or false.

The difference between float32 and float64 is the precision, meaning that float64 will represent the number with higher accuracy, but will take more space in memory.

```go
var x int = 42
var y float32 = 1.37
var name string = "James"
var online bool = true
```

Since we used initializers when declaring the variables, we could have omitted the types, as they would be taken from the values assigned. We included them to demonstrate the different data types.

Another interesting feature of Go are **zero values**: variables that are declared without a value take the zero value of their type. `0` for numeric types, `false` for the boolean type, `""` for strings.

## Constants

A variable can change its value during the program but in some cases your program may need values that are preserved during the program. These are called constants and they cannot be changed from their initial value.

Constants are declared like variables, but with the `const` keyword and need to be assigned a value:

```go
const pi = 3.14 
fmt.Println(pi)
```

Now, `pi` is a constant and cannot be changed. Constants cannot be declared using the `:=` syntax.

## Input

The `"fmt"` package also allows you to take input from the user of the program.
To take input, we need to use the `Scanln()` function and provide it with the variable which should hold the input value:

```go
var input string
fmt.Scanln(&input)
fmt.Println(input)
```

Now the input variable will hold the value which the user enters when running the program.

*Note the ampersand & before the variable name -- it is used to return the address of a variable.*

The previous example took the input from the user as a string.
If we need to take a number as input, such as an integer, we can simply declare the type of the input variable as int and Go will automatically convert the input to that type:

```go
var input int
fmt.Scanln(&input)
fmt.Println(input*2) 
```
