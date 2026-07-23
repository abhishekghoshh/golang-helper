# Generics

Generics allow our functions or data structures to take in several types that are defined in their generic form.

To truly understand what this means, let's take a look at a very simple case.

Let's say you need to make a function that takes one slice and prints it. Then you might write this type of function:

```go
func Print(s []string) {
    for _, v := range s {
        fmt.Print(v)
    }
}
```

Simple, right? What if we want to have the slice be an integer? You will need to make a new method for that:

```go
func Print(s []int) {
    for _, v := range s {
        fmt.Print(v)
    }
}
```

These solutions might seem redundant, as we're only changing the parameter. But currently, that's how we solve it in Go without resorting to making it into some interface.

And now with generics, they will allow us to declare our functions like this:

```go
func Print[T any](s []T) {
    for _, v := range s {
        fmt.Print(v)
    }
}
```

**In the above function, we are declaring two things:**

- We have T, which is the type of the any keyword (this keyword is specifically defined as part of a generic, which indicates any type)
- And our parameter, where we have variable s whose type is a slice of T.

We will now be able to call our method like this:

```go
func main() {
    Print([]string{"Hello, ", "playground\n"})
    Print([]int{1,2,3})
}
```

## Generic Constraints

There are limitations on how far generics can take us. Printing, for example, is pretty simple since Golang can print out any type of variable being thrown into it.

What if we want to do more complex things? Let's say that we have defined our own methods for a structure and want to call it:

```go
type worker string

func (w worker) Work(){
    fmt.Printf("%s is working\n", w)
}
func DoWork[T any](things []T) {
    for _, v := range things {
        v.Work()
    }
}

func main() {
    var a,b,c worker
    a = "A"
    b = "B"
    c = "C"
    DoWork([]worker{a,b,c})	
}
```

And we will get this:

```
type checking failed for main
prog.go2:25:11: v.Work undefined (type bound for T has no method Work)
```

It fails to run because the slice processed inside the function is of type any and it doesn't implement the method Work, which makes it fail to run.

We can actually make it work, though, by using an interface:

```go
type Person interface {
    Work()
}

type worker string

func (w worker) Work(){
    fmt.Printf("%s is working\n", w)
}

func DoWork[T Person](things []T) {
    for _, v := range things {
        v.Work()
    }
}
func DoWorkInterface(things []Person) {
    for _, v := range things {
        v.Work()
    }
}
func main() {
    var a,b,c worker
    a = "A"
    b = "B"
    c = "C"
    DoWork([]worker{a,b,c})

    var d,e,f worker
    d = "D"
    e = "E"
    f = "F"
    DoWorkInterface([]Person{d,e,f})
}
```

And it will print out:

```
A is working
B is working
C is working
```
