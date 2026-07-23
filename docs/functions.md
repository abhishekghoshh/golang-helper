# Functions

Functions are an important part of Go. They allow you to define a block of code, give it a name and call it in the code. This makes the function code reusable, as you can call it multiple times in different parts of your program.

We have already seen functions in our previous examples. For example, the Println() function outputs text. Println is the name of the function, while "Hello" is the argument which we provide to the function using the parentheses. Remember, to call a function, we always use the name of the function followed by parentheses.

Function arguments are a way for our functions to take input parameters. For example, we can add a name argument to our welcome() function and use it in the output.

As you can see, the argument is provided in the parentheses and includes the name followed by the type. The argument behaves like a variable inside the function's body, meaning you can use its value by its name. When calling the function, we need to pass it a string argument inside the parentheses.

We are able to reuse the code in our function by calling it with different arguments. Remember, when defining the function, the type of the argument comes after the variable name.

To make a function take multiple arguments, separate them using commas. For example, let's create a function that takes two integer arguments and outputs their sum.

If the arguments are of the same type, you can omit the types for the arguments and define it only for the last one.

All of our function examples so far have resulted in a simple output. Oftentimes, the result of the function needs to be assigned to a variable so that your code can use it further.

The return statement terminates the function and returns the provided value to the code that called it. Notice that we need to define the return type of the function after the arguments definition. In our case, it is int.

```go
package main
import "fmt"
func main() {
  fmt.Println("Hello") 
  result := sum( 1 , 3 )
  fmt.Println( result )
}
func sum(a int, b int) int {
  fmt.Println(a+b)
  return a+b
}
func sum(a, b int){
   fmt.Println(a+b)
}
```

```go
package main
import "fmt"
func welcome(name string) {
   fmt.Println("Welcome", name)
}
func main() {
    welcome("David")
    welcome("James")
}
```

## Multiple Return Values

A useful feature of Go is that functions can return multiple values. The swap() function takes two integer arguments and returns them in swapped order.
Note that the return type of each value should be declared: in our case it is (int, int). We can call our function and assign it result to variables:

```go
func swap(x, y int) (int, int) {
    return y, x
}

func main() {
    a, b := swap(42, 8)
    fmt.Println(a)
    fmt.Println(b)
}
```

Returning multiple values from a function is handy! For example, it can be used to return both the result and the error values of an operation.

## Defer

A `Defer` statement ensures that the function is called only after the surrounding function returns.

```go
func welcome() {
    fmt.Println("Welcome")
}

func main() {
    defer welcome()
    fmt.Println("Hey")
}
```

The code will first output "Hey" and only after that output the result of the welcome() function. This happens because the call to welcome() is deferred, meaning it waits until `main()` finishes execution and only then calls it.

defer is often used for clean-up, for example, to release resources used by the code, such as files, connections, etc.

If you have deferred multiple function calls, they will execute in **last-in-first-out** order. The defer calls are stacked on top of each other, which is why they are executed in last-in-first-out order.

```go
func main() {
    fmt.Println("start")

    for i := 0; i < 5; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("end")
}
```

## Scope

Scope is where a variable can be used. There are two main scopes in Go: local and global.

A variable defined in the function is called a **local variable**. Their scope is only in the function body, which means they only exist within their function:

```go
func test() {
  var x = 8
}
```

Now, x is a local variable and is available only in the body of the test() function. Trying to access it in another function, for example main(), will cause an error.

A variable defined outside the local scope is called a **global variable**. Global variables can be used throughout the package:

```go
var x = 8
func test() {
  fmt.Println(x)
}
func main() {
  fmt.Println(x)
}
```

The variable x is declared outside of the functions, making it a global variable, which is accessible anywhere in the package.

Global variables are often considered a bad practice. It is better to pass variables as function arguments.

## Variadic Functions

Variadic functions accept a variable number of trailing arguments. Use `func(args ...Type)` syntax and pass slices with `slice...`.

```go
func sum(nums ...int) {
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
    sum(1, 2)           // individual args
    sum(1, 2, 3)
    nums := []int{1, 2, 3, 4}
    sum(nums...)        // spread a slice
}
```

## Recursion

A function that calls itself. Closures can be recursive but must be declared with a typed `var` before definition.

```go
func fact(n int) int {
    if n == 1 { return n }
    return n * fact(n-1)
}

func main() {
    fmt.Println("factorial of 7 is", fact(7)) // 5040

    var fib func(n int) int
    fib = func(n int) int {
        if n < 2 { return n }
        return fib(n-1) + fib(n-2)
    }
    fmt.Println("7th fibonacci is", fib(7)) // 13
}
```

## Closures

A closure is a function that references variables from its enclosing scope. Each closure captures its **own independent state**.

```go
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    nextInt1 := intSeq()
    fmt.Println(nextInt1()) // 1
    fmt.Println(nextInt1()) // 2

    nextInt2 := intSeq()    // new instance, independent state
    fmt.Println(nextInt2()) // 1
}
```

## Panic & Recover

`panic` stops normal execution and begins unwinding. `recover` (called inside a deferred function) catches a panic. Use sparingly — Go favors explicit error returns.

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    panic("This is a custom panic")
    fmt.Println("Never reaches here")
}
```

## Custom Default Parameters

Go has no built-in default function parameters. Two workarounds:

**1. Variadic arguments:**
```go
func greetWithDefault(names ...string) string {
    if len(names) == 0 {
        return greet("Abhishek Ghosh") // default
    }
    return greet(names[0])
}
```

**2. Functional Options Pattern:**
```go
type GreetingOptions struct { Name string; Age int }
type GreetingOption func(*GreetingOptions)

func WithName(name string) GreetingOption {
    return func(o *GreetingOptions) { o.Name = name }
}
func WithAge(age int) GreetingOption {
    return func(o *GreetingOptions) { o.Age = age }
}

func GreetWithDefaultOptions(options ...GreetingOption) string {
    opts := GreetingOptions{Name: "Abhishek Ghosh", Age: 25}
    for _, o := range options { o(&opts) }
    return Greet(opts)
}

// Usage: GreetWithDefaultOptions(WithName("Alice"), WithAge(20))
```

## Custom Types with Methods

Methods can be declared on named types (not just structs).

```go
type color string

func (c color) describe(description string) string {
    return string(c) + " " + description
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 { return float64(-f) }
    return float64(f)
}

func main() {
    c := color("Red")
    fmt.Println(c.describe("is an awesome color"))

    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs()) // 1.414...
}
```
