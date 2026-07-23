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

## Complete Concurrency Examples

### Goroutines
```go
func GoRoutines() {
    var out = func(from, to int) {
        for i := from; i <= to; i++ {
            fmt.Print(i, " ")
        }
    }

    go out(0, 5)
    go out(6, 10)
    time.Sleep(time.Second) // wait for goroutines
}
```

### Channel Buffering
Buffered channels accept values without a corresponding receiver (up to capacity). Also useful for limiting concurrent goroutines.

```go
// Buffered channel
messages := make(chan string, 2)
messages <- "buffered"   // doesn't block
messages <- "channel"    // doesn't block
fmt.Println(<-messages)

// Using buffered channels as a semaphore to limit goroutines
eventLimiter := make(chan interface{}, 5)
for event := range eventManager.Stream() {
    eventLimiter <- true        // blocks if 5 goroutines are running
    go func() {
        process(event)
        <-eventLimiter           // release slot
    }()
}
```

### Channel Directions
Specify send-only (`chan<-`) or receive-only (`<-chan`) to enforce channel usage at compile time.

```go
var ping = func(pings chan<- string, names []string) {
    for _, name := range names {
        pings <- "Hi ping," + name
    }
    close(pings)
}

var pong = func(pings <-chan string, pongs chan<- string) {
    for name := range pings {
        pongs <- name + ", Hi pong"
    }
    close(pongs)
}
```

### Channel Synchronization
Use channels to synchronize goroutines without `WaitGroup`.

```go
var worker = func(n int, done chan bool) {
    for i := 0; i < n; i++ {
        fmt.Print(i, " ")
    }
    done <- true
}

done := make(chan bool, 1)
go worker(5, done)
<-done // block until worker signals completion
```

### Non-blocking Channel Operations
Use `select` with `default` for non-blocking sends/receives.

```go
messages := make(chan string)

// Non-blocking receive
select {
case msg := <-messages: fmt.Println("received", msg)
default:                fmt.Println("no message received")
}

// Non-blocking send
select {
case messages <- "hi": fmt.Println("sent")
default:               fmt.Println("no message sent")
}
```

### Closing Channels
Closed channels return zero value immediately. Use `val, more := <-ch` to detect closure.

```go
jobs := make(chan int, 5)
done := make(chan bool)

go func() {
    for {
        j, more := <-jobs
        if more { fmt.Println("received job", j) }
        else    { done <- true; return }
    }
}()

for j := 1; j <= 6; j++ { jobs <- j }
close(jobs)
<-done
```

### Range Over Channels
```go
queue := make(chan string, 2)
queue <- "one"; queue <- "two"
close(queue)

for elem := range queue {
    fmt.Println(elem) // "one", "two"
}
```

### Timers
```go
timer1 := time.NewTimer(2 * time.Second)
<-timer1.C
fmt.Println("Timer 1 fired")

// Cancellable timer
timer2 := time.NewTimer(time.Second)
stop2 := timer2.Stop()
if stop2 { fmt.Println("Timer 2 stopped") }
```

### Tickers
```go
ticker := time.NewTicker(500 * time.Millisecond)
done := make(chan bool)

go func() {
    for {
        select {
        case <-done:       return
        case t := <-ticker.C: fmt.Println("Tick at", t)
        }
    }
}()

time.Sleep(1600 * time.Millisecond)
ticker.Stop()
done <- true
```

### Worker Pools
```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= numJobs; j++ { jobs <- j }
    close(jobs)

    for a := 1; a <= numJobs; a++ { <-results }
}
```

### Wait Groups
```go
var wg sync.WaitGroup
var worker = func(id int) {
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

for i := 1; i <= 5; i++ {
    wg.Add(1)
    go func(num int) {
        defer wg.Done()
        worker(num)
    }(i)
}
wg.Wait()

// Producer-consumer with WaitGroup
ch := make(chan int, 1)
consumer := func(id int) {
    defer wg.Done()
    for i := range ch {
        fmt.Printf("Reader %d: %d\n", id, i)
    }
}
for i := 1; i <= 5; i++ {
    wg.Add(1)
    go consumer(i)
}
for i := 0; i < 100; i++ { ch <- i }
close(ch)
wg.Wait()
```

### Rate Limiting
```go
requests := make(chan int, 5)
for i := 1; i <= 5; i++ { requests <- i }
close(requests)

// Simple: 1 request per 200ms
limiter := time.Tick(200 * time.Millisecond)
for req := range requests {
    <-limiter
    fmt.Println("request", req, time.Now())
}

// Bursty: buffer of 3, refilled every 200ms
burstyLimiter := make(chan time.Time, 3)
for i := 0; i < 3; i++ { burstyLimiter <- time.Now() }
go func() {
    for t := range time.Tick(200 * time.Millisecond) {
        burstyLimiter <- t
    }
}()
```

### Atomic Counters
```go
var ops atomic.Uint64
var wg sync.WaitGroup

for i := 0; i < 50; i++ {
    wg.Add(1)
    go func() {
        for c := 0; c < 1000; c++ {
            ops.Add(1) // thread-safe increment
        }
        wg.Done()
    }()
}
wg.Wait()
fmt.Println("ops:", ops.Load()) // exactly 50000
```

### Mutexes
For complex state, use `sync.Mutex`.

```go
type Container struct {
    mu       sync.Mutex
    counters map[string]int
}

func (c *Container) inc(name string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.counters[name]++
}
```

### Stateful Goroutines (Channel-based state)
Alternative to mutexes: a single goroutine owns the state, others communicate via channels.

```go
type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

// State-owning goroutine
go func() {
    var state = make(map[int]int)
    for {
        select {
        case read := <-reads:
            read.resp <- state[read.key]
        case write := <-writes:
            state[write.key] = write.val
            write.resp <- true
        }
    }
}()
```
