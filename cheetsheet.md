
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
Now, `pi` is a constant and cannot be changed. Constants cannot be declared using the `:=` syntax. </br>


### Input

The `"fmt"` package also allows you to take input from the user of the program.</br>
To take input, we need to use the `Scanln()` function and provide it with the variable which should hold the input value: </br>

```
var input string
fmt.Scanln(&input)
fmt.Println(input)
```
Now the input variable will hold the value which the user enters when running the program.</br>
*Note the ampersand & before the variable name -- it is used to return the address of a variable.* </br>
The previous example took the input from the user as a string. </br>
If we need to take a number as input, such as an integer, we can simply declare the type of the input variable as int and Go will automatically convert the input to that type:</br>
```
var input int
fmt.Scanln(&input)
fmt.Println(input*2) 
```


### Decision Making

The if statement can be used to make decisions and run code when a given condition is true. For example: </br>
```
x := 42
if x > 18 {
    fmt.Println("Allowed")
}
```
The code of the if statement should be enclosed in curly braces { } and can include multiple lines of code. </br>
The code in the curly braces of an if statement will run only if the condition evaluates to true. </br>
The else statement can be used to run code when the condition of an if statement is false. For example: </br>
```
x := 14
if x > 18 {
    fmt.Println("Allowed")
} else {
    fmt.Println("Not allowed")
}
```
There is no ternary if in Go, so you'll need to use a full if statement even for basic conditions.</br>
Sometimes you need a variable only for the if/else statements.</br>
For this, the if statement in Go can start with a variable declaration before the condition:</br>
```
if x := 42; x > 18 {
    fmt.Println("Allowed")
} else {
    fmt.Println("Not allowed")
}
```
Note the semicolon after the variable declaration - it is used to separate the statements.</br>
Variables declared in a if statement are only available in the if/else blocks.</br>

A `switch` statement is a shorter way of writing a sequence of if/else statements.</br>
For example, imagine having the following if/else statements. The code checks the value of the num variable and outputs the corresponding text version. </br>
We can achieve the same result using a switch statement.</br>
```
num := 3
if num == 1 {
    fmt.Println("One")
} else if num == 2 {
    fmt.Println("Two")
} else if num == 3 {
    fmt.Println("Three")
} else {
    fmt.Println(num)
}
```
```
num := 3
switch num {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
    case 3:
        fmt.Println("Three")
    default:
        fmt.Println(num)
}
```
The switch statement makes the code shorter and easier to read. Each case statement includes the value to compare with and the code to run after the colon :</br>
A simple way to replace an if/else chain is to use a switch without an expression. That way, each case statement can include a condition:</br>
```
x := 2
switch {
    case x>0 && x<10:
        fmt.Println("something")
    case x > 10:
        fmt.Println("something else")
}
```
The optional default case runs if none of the cases match.</br>
In other programming languages, such as C++ or Java, each case needs to have a break statement, in order to stop the execution of the switch statement. </br>
In Go, the break statement is not needed, as the switch cases evaluate from top to bottom, stopping when a case succeeds.</br>
Similar to if statements, switch can also have a short variable declaration before the conditional expression.</br>


## Loop

Loops are used to repeat code until a certain condition holds. </br>
The only loop construct in Go is for, which has three components: init, condition, post statement. </br>
For example:
```
for i:=0; i<5; i++ {
  fmt.Println(i)
}
```
The code above will output the numbers 0 to 4, as we initialize the value of i to 0 and increment it while the condition `i<5` holds. </br>
`i++` is a shorter version of `i = i+1` and is called the increment operator. </br>
The init and post statements can be omitted.</br>
```
num, sum := 1, 0
for num <= 1000 {
   sum += num
   num++
}
fmt.Println(sum)
```
The code above will calculate the sum of numbers from 1 to 1000. Note that with just the condition, for becomes similar to while loops available in other programming languages.</br>


## Functions

Functions are an important part of Go. They allow you to define a block of code, give it a name and call it in the code. This makes the function code reusable, as you can call it multiple times in different parts of your program.
We have already seen functions in our previous examples. For example, the Println() function outputs text. Println is the name of the function, while "Hello" is the argument which we provide to the function using the parentheses. Remember, to call a function, we always use the name of the function followed by parentheses.
Function arguments are a way for our functions to take input parameters. For example, we can add a name argument to our welcome() function and use it in the output.
As you can see, the argument is provided in the parentheses and includes the name followed by the type. The argument behaves like a variable inside the function's body, meaning you can use its value by its name. When calling the function, we need to pass it a string argument inside the parentheses.
We are able to reuse the code in our function by calling it with different arguments. Remember, when defining the function, the type of the argument comes after the variable name.
To make a function take multiple arguments, separate them using commas. For example, let's create a function that takes two integer arguments and outputs their sum.
If the arguments are of the same type, you can omit the types for the arguments and define it only for the last one.
All of our function examples so far have resulted in a simple output. Oftentimes, the result of the function needs to be assigned to a variable so that your code can use it further.
The return statement terminates the function and returns the provided value to the code that called it. Notice that we need to define the return type of the function after the arguments definition. In our case, it is int.
```
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