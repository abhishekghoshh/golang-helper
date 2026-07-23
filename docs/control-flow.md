# Control Flow

## Decision Making

### If / Else

The if statement can be used to make decisions and run code when a given condition is true. For example:

```go
x := 42
if x > 18 {
    fmt.Println("Allowed")
}
```

The code of the if statement should be enclosed in curly braces { } and can include multiple lines of code.

The code in the curly braces of an if statement will run only if the condition evaluates to true.

The else statement can be used to run code when the condition of an if statement is false. For example:

```go
x := 14
if x > 18 {
    fmt.Println("Allowed")
} else {
    fmt.Println("Not allowed")
}
```

There is no ternary if in Go, so you'll need to use a full if statement even for basic conditions.

Sometimes you need a variable only for the if/else statements.
For this, the if statement in Go can start with a variable declaration before the condition:

```go
if x := 42; x > 18 {
    fmt.Println("Allowed")
} else {
    fmt.Println("Not allowed")
}
```

Note the semicolon after the variable declaration - it is used to separate the statements.
Variables declared in an if statement are only available in the if/else blocks.

### Switch

A `switch` statement is a shorter way of writing a sequence of if/else statements.
For example, imagine having the following if/else statements. The code checks the value of the num variable and outputs the corresponding text version.
We can achieve the same result using a switch statement.

```go
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

```go
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

The switch statement makes the code shorter and easier to read. Each case statement includes the value to compare with and the code to run after the colon `:`.

A simple way to replace an if/else chain is to use a switch without an expression. That way, each case statement can include a condition:

```go
x := 2
switch {
    case x>0 && x<10:
        fmt.Println("something")
    case x > 10:
        fmt.Println("something else")
}
```

The optional default case runs if none of the cases match.

In other programming languages, such as C++ or Java, each case needs to have a break statement, in order to stop the execution of the switch statement.
In Go, the break statement is not needed, as the switch cases evaluate from top to bottom, stopping when a case succeeds.

Similar to if statements, switch can also have a short variable declaration before the conditional expression.

## Loops

Loops are used to repeat code until a certain condition holds.
The only loop construct in Go is `for`, which has three components: init, condition, post statement.
For example:

```go
for i:=0; i<5; i++ {
  fmt.Println(i)
}
```

The code above will output the numbers 0 to 4, as we initialize the value of i to 0 and increment it while the condition `i<5` holds.
`i++` is a shorter version of `i = i+1` and is called the increment operator.

The init and post statements can be omitted.

```go
num, sum := 1, 0
for num <= 1000 {
   sum += num
   num++
}
fmt.Println(sum)
```

The code above will calculate the sum of numbers from 1 to 1000. Note that with just the condition, for becomes similar to while loops available in other programming languages.

## Complete Example

```go
func FlowControl() {
    // Standard for loop
    for i := 1; i <= 10; i++ {
        fmt.Print(i, " ")
    }

    // while-style loop (omit init and post)
    sum := 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println("sum is", sum)

    // Infinite loop with break
    num := 0
    for {
        if num == 10 { break }
        num++
    }

    // if with short statement
    fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

    // Switch on OS
    switch os := runtime.GOOS; os {
    case "darwin": fmt.Println("OS X.")
    case "linux":  fmt.Println("Linux.")
    default:       fmt.Printf("%s.\n", os)
    }

    // Switch without expression (clean if-else chains)
    t := time.Now()
    switch {
    case t.Hour() < 12: fmt.Println("Good morning!")
    case t.Hour() < 17: fmt.Println("Good afternoon.")
    default:            fmt.Println("Good evening.")
    }

    // Type switch
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:   fmt.Println("I'm a bool")
        case int:    fmt.Println("I'm an int")
        default:     fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}
```
