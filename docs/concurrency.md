# Concurrency

Concurrency means multiple computations are happening at the same time. It is used when your program has multiple things to do. Concurrency is about creating multiple processes executing independently.

For example, imagine a shooting game, where many things are happening at the same time: enemies are running and shooting, points are being calculated, weapons are being unlocked, etc.
All of these things need to be happening concurrently.

In order to use concurrency, the program is broken into parts, which are then executed separately.
When using concurrency, we are able to achieve the intended results in less time, thus increasing the overall performance and efficiency of our programs.

To achieve concurrency, Go provides **Goroutines**.

## Goroutines

A Goroutine is much like a thread to accomplish multiple tasks, but it consumes fewer resources than OS threads.
When a program is broken down into separate tasks, each Goroutine can be used to accomplish a task, enabling concurrency in our program.

Goroutines are not OS threads, they are virtual threads, managed by Go. You can have thousands of Goroutines running in a Go program. Let's have a look at the following program:

```go
import (
    "fmt"
    "time"
)
func out(from, to int) {
    for i := from; i <= to; i++ {
        time.Sleep(50 * time.Millisecond)
        fmt.Println(i)
    }
}
func main() {
    out(0, 5)
    out(6, 10)
}
```

The `out()` function simply outputs numbers in the given range. We use a `time.Sleep()` to emulate work being done between the outputs just for demonstration purposes. It simply waits for the provided time (50ms) before continuing the execution.

Now, if we call the function in main two times, the first call will execute first followed by the second call.
This will generate the output of 0 to 5, then 6 to 10.

This is called a sequential program, as the statements are executed one after the other. The first call needs to complete, before the second call starts.
When running concurrent programs, we often do not want to wait for one task to finish before starting a new one.

To achieve concurrency, let's start the function calls as Goroutines, using the `go` keyword:

```go
go out(0, 5)
go out(6, 10)
```

It is very simple to start a Goroutine -- we simply need to add a `go` keyword before the function call. If we run the program, we get No Output. This output happens because the `main()` function exits before the Goroutines complete. Our program has 3 virtual threads. The 2 function calls, and `main()`. Our 2 function calls get executed concurrently, and `main()` does not wait for them to finish.

```go
go out(0, 5)
go out(6, 10)
time.Sleep(500 * time.Millisecond)
```

The 500ms wait should be enough time for the Goroutines to finish executing and generating the output. Now when you run the code, you will see that the output is not sequential, each Goroutine worked independently and concurrently.

## Channels

Goroutines run independently and they do not know when another one has finished executing. This causes, for example, the `main()` function to quit, before any started Goroutine has finished. To enable communication between Goroutines, Go provides **Channels**.

A channel is like a pipe, allowing you to send and receive data from Goroutines, and enabling them to communicate and synchronize.
Similar to how water flows from one end to another in a pipe, data can be sent from one end and received by the other end using channels. To use a channel, we first need to make one using the `make()` function.

```go
ch := make(chan int)
```

The type after the `chan` keyword indicates the type of the data we will send through the channel.

We can send data to the channel using the following syntax:

```go
ch <- 8
```

Similarly, we can receive data from the channel using the following syntax:

```go
value := <- ch
```

If we do not need the value as a variable, we can simply use:

```go
<- ch
```

Data flows in the direction of the arrow.

We can use our channel and rewrite the previous example without a `time.Sleep()` in `main()`:

```go
import (
    "fmt"
    "time"
)

func out(from, to int, ch chan bool) {
    for i := from; i <= to; i++ {
        time.Sleep(50 * time.Millisecond)
        fmt.Println(i)
    }
    ch <- true
}

func main() {
    ch1 := make(chan bool)
    ch2 := make(chan bool)

    go out(0, 5, ch1)
    go out(6, 10, ch2)

    fmt.Println(<-ch1)
    fmt.Println(<-ch2)
}
```

We define a bool channel and pass it to our `out()` function as an argument. After the function finishes its task, we send the value `true` to the channel, which is received in `main()`.

Now, `main()` is waiting to receive data from the channel, making the program wait for the Goroutines to finish executing.

The receive operation blocks the code until and unless some data is sent by the send operation. If no data is received, a deadlock will occur, blocking the code from executing.

Let's use a channel to send data from a Goroutine and use it in `main()`.
Our program needs to calculate and output the sum of all even numbers in a given range plus the sum of their squares and output the result:

```
output = evenSum + squareSum
```

We will use two Goroutines: one to calculate the evenSum, and the other to calculate the squareSum. We will get the data using channels in `main()`, then calculate and output the final sum.

```go
func evenSum(from, to int, ch chan int) {
    result := 0
    for i := from; i <= to; i++ {
        if i%2 == 0 {
            result += i
        }
    }
    ch <- result
}
func squareSum(from, to int, ch chan int) {
    result := 0
    for i := from; i <= to; i++ {
        if i%2 == 0 {
            result += i * i
        }
    }
    ch <- result
}

func main() {
    evenCh := make(chan int)
    sqCh := make(chan int)

    go evenSum(0, 100, evenCh)
    go squareSum(0, 100, sqCh)

    fmt.Println(<-evenCh + <-sqCh)
}
```

As you can see, our functions send the result via channels.
Now we can call them as Goroutines in `main()` and output the resulting sum.
We use the channels to get the result of each Goroutine and output their sum.

If you do not need to send data to a channel anymore, you can close it using `close(ch)`, where ch is the name of the channel. This is done in the sender.

## Select

The `select` statement is used to wait on multiple channel operations.
The syntax is similar to switch except that each of the case statements will be a channel operation.

Let's use the same program from our previous example and select the channel that is ready first:

```go
func evenSum(from, to int, ch chan int) {
    result := 0
    for i := from; i <= to; i++ {
        if i%2 == 0 {
            result += i
        }
    }
    ch <- result
}
func squareSum(from, to int, ch chan int) {
    result := 0
    for i := from; i <= to; i++ {
        if i%2 == 0 {
            result += i * i
        }
    }
    ch <- result
}

func main() {
    evenCh := make(chan int)
    sqCh := make(chan int)

    go evenSum(0, 100, evenCh)
    go squareSum(0, 100, sqCh)

    select {
        case x := <-evenCh:
            fmt.Println(x)
        case y := <-sqCh:
            fmt.Println(y)
    }
}
```

The `select` statement waits for a channel to receive data and executes its case. This means that only one of the cases will execute -- the one that corresponds to the channel that receives data first. If both channels receive data at the same time, one of the cases is chosen randomly.

Combining Goroutines and channels with select is a powerful feature of Go. Imagine a program that needs to execute some code whenever one of the concurrent operations completes -- this can be achieved using select.

A select can have a `default` case, which will execute when no channel is ready. For example, we could have an infinite for loop, waiting for one of the channels to receive data:

```go
select {
    case x := <-evenCh:
        fmt.Println(x)
    case y := <-sqCh:
        fmt.Println(y)
    default:
        fmt.Println("This is default case")
    }
```

The for loop uses a select to check which channel got data. If none of them are ready, the default case will execute which will wait for 50ms.
As soon as a channel gets data, the return statement will exit the loop.

A `select` statement blocks until at least one of its cases can proceed. The default case is useful in preventing deadlocks -- without it the select would wait for a channel forever, crashing the program if none of the channels received data.

## Learning Resources

### Youtube

- [Go Concurrency](https://www.youtube.com/playlist?list=PL7g1jYj15RUNqJStuwE9SCmeOKpgxC0HP)

### Udemy

- [Concurrency in Go (Golang)](https://www.udemy.com/course/concurrency-in-go-golang)
  - [go-concurrency-exercises](https://github.com/andcloudio/go-concurrency-exercises)
- [Working with Concurrency in Go (Golang)](https://www.udemy.com/course/working-with-concurrency-in-go-golang/)
