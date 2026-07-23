# Pointers

All of the values that we define in our program are stored in the computer memory and have their own unique memory address.
Pointers are special variables that hold the memory address of values.

In Go, we declare a pointer using a `*`:
`var p *int` — Now, p is a pointer to an integer value.

We know how to define a pointer, but how do we assign it a memory address? This is done using the `&` operator, which returns the memory address of a variable.

```go
x := 42
p := &x
```

Now `p` is a pointer and holds the memory address of `x`.

If we want to access the underlying value of a pointer, we can use the `*` operator:

```go
x := 42
p := &x
fmt.Println(*p)
```

The `*` operator can also be used to change the value of the memory address the pointer holds:

```go
x := 42
p := &x
*p = 8
fmt.Println(*p)
fmt.Println(x)
```

The `*` operator is called the **dereferencing operator**.

We have used pointers with functions in a previous lesson, when taking input from the user:

```go
var input string
fmt.Scanln(&input)
fmt.Println(input)
```

Here, we pass the memory address of the input variable (a pointer to input) to the `Scanln()` function, which uses it to store the input value.

## Pointers as Function Parameters

```go
func change(val int) {
  val = 8
}
func change_ptr(ptr *int) {
  *ptr = 8
}
func main() {
  x := 42
  change(x)
  fmt.Println(x)
  change_ptr(&x)
  fmt.Println(x)
}
```

The `change()` function takes an integer argument and changes its value. The `change_ptr()` function does the same using a pointer.

When you run the code, you will see that the change() function did not change the value of our x variable, because the argument is just a copy of its value, while the `change_ptr()` did change the actual value of `x`, because it used its memory address as the argument.

Note that we need to pass the memory address using the & operator to functions that take a pointer as their argument.
