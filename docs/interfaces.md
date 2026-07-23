# Interfaces

In Go programming, we use interfaces to store a set of methods without implementation. That is, methods of interface won't have a method body. For example:

```go
type Shape interface {
  area() float32
  perimeter() float32
}
```

Here, `Shape` is an interface with methods: `area()` and `perimeter()`. You can see both methods only have method signatures without any implementation.

To use an interface, we first need to implement it by a type (struct). To implement an interface, a struct should provide implementations for all methods of an interface. For example:

```go
package main 
import "fmt"

type Shape interface {
  area() float32
}

type Rectangle struct {
  length, breadth float32
}

func (r Rectangle) area() float32 {
  return r.length * r.breadth
}

func calculate(s Shape) {
  fmt.Println("Area:", s.area())
}

func main() {
  rect := Rectangle{7, 4}
  calculate(rect)
}
```

When a struct implements an interface, it should provide an implementation for all the methods of the interface. If it fails to implement any method, we will get an error.

```
cannot use r (type Rectangle) as type Shape in argument to calculate:
    Rectangle does not implement Shape (missing perimeter method)
```

## Empty Interface

We know that interfaces are used to store a set of methods without implementation. In Go, we can also create interfaces without any methods, known as **empty interfaces**. For example: `interface {}`

Here, we have created an empty interface without any methods.
In Go, we can create variables of the empty interface type.

```go
package main
import "fmt"

func main() {
  var a interface {}
  fmt.Println("Value:", a)
}
```

Here, we have created a variable of type empty interface. When we print the variable, we get nil as output.

### Empty Interface to Pass Any Type

Normally in functions, when we pass values to the function parameters, we need to specify the data type of parameters in a function definition.
However, with an empty interface, we can pass parameters of any data type.
We can also use an empty interface to pass any number of arguments to the function definition.

```go
package main
import "fmt"

func displayValue(i interface {}) {
  fmt.Println(i)
}
func displayValues(i... interface {}) {
  fmt.Println(i)
}

func main() {
  a := "Welcome to Programiz"
  b := 20
  c := false
  displayValue(a)
  displayValue(b)
  displayValue(c)

  displayValues(a)
  displayValues(a, b)
  displayValues(a, b, c)
}
```

In the above example, we have used an empty interface `i` as the parameter to the `displayValue()` function.
Now the same function works for any type of parameters (string, numeric, boolean).

## Type Assertions

Type assertions allow us to access the data and data type of values stored by the interface.

```go
var a interface{}
a = "Hello World"
a = 10
```

Here, the `a` variable is of empty interface type, and it is storing both the string and integer value.
This seems like an important feature of an empty interface. However, sometimes this will create ambiguity on what data an interface is holding.
To remove this ambiguity, we use type assertions.

```go
package main
import "fmt"

func main() {
  // create an empty interface
  var a interface {}

  // store integer to an empty interface
  a = 10

  // type assertion
  interfaceValue := a.(int)

  fmt.Println(interfaceValue)
}
```

In the above example, we have stored the integer value 10 to the empty interface denoted by a. Notice the code:

Here, `(int)` checks if the value of a is an integer or not. If true, the value will be assigned to `interfaceValue`.
Otherwise, the program panics and terminates.

The type assertion statement actually returns a boolean value along with the interface value:

```go
var a interface {}
a = 12
interfaceValue := a.(int)
```

Here, the data type of value 12 matches with the specified type `(int)`, so this code assigns the value of a to `interfaceValue`.
Along with this, a.(int) also returns a boolean value which we can store in another variable:

```go
interfaceValue, booleanValue := a.(int)
```

We get both the data and boolean value.
We can use this to avoid panic if the data type doesn't match the specified type.
This is because, when the type doesn't match, it returns false as the boolean value, so the panic doesn't occur.

Output:
```
Interface Value: 0
Boolean Value: false
```
