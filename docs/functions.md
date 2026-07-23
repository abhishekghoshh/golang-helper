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
