
### Basics

Sample GO code
```
package main // Package declaration

import "fmt" // Import other packages that we need

func main() {  //  Declare functions, tell Go to do things
    fmt.Println("Hello World")
}
```

To import a package, we use the `import` statement
One of the most popular packages is `"fmt"`, which stands for format, and provides input and output functionality.</br>
Package `main` is the entry point for go program. </br>
You can import multiple packages at once, using parentheses. For example: `import ("fmt" "math")` </br>

In Go, the var keyword is used to declare variables.</br>
For example: `var i int`</br>
```
var i int = 8
fmt.Println(i) 
```


### Variables

The code above declares a variable named i of type int. int stands for integer and represents whole numbers. </br>
We can assign the variable a value and output it.</br>
You can also declare multiple variables on one line and assign values to them</br>
```
var i, j int = 8, 42
fmt.Println(i) 
fmt.Println(j)
```
If you assign a value to the variable, you can omit the type declaration, as Go will automatically take the type of the value assigned.</br>
Go supports short variable declaration using `:=` </br>
For example :` i := 8` </br>
The variable `i` will be a integer variable and the value of the `i` will be `8` </br>
It can also be used to declare and initialize multiple variables on one line. ` x, y := 10, 5` </br>


### Data Types

We used int to declare integer values in our previous lesson. `var i int = 8 ` </br>
Let's see what other common types Go supports.</br>
**float32** - a single-precision floating point value.</br>
**float64** - a double-precision floating point value.</br>
**string** - a text value.</br>
**bool** - Boolean true or false.</br>
The difference between float32 and float64 is the precision, meaning that float64 will represent the number with higher accuracy, but will take more space in memory.</br>
```
var x int = 42
var y float32 = 1.37
var name string = "James"
var online bool = true
```

Since we used initializers when declaring the variables, we could have omitted the types, as they would be taken from the values assigned. We included them to demonstrate the different data types.</br>
Another interesting feature of Go are zero values: variables that are declared without a value take the zero value of their type. 0 for numeric types, false for the boolean type,"" for strings.</br>



### Constants

A variable can change its value during the program but in some cases your program may need values that are preserved during the program. These are called constants and they cannot be changed from their initial value.</br>
Constants are declared like variables, but with the `const` keyword and need to be assigned a value:</br>
```
const pi = 3.14 
fmt.Println(pi)
```
