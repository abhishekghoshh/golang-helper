# Others

Miscellaneous Go examples.

## Go Curl

Simple command-line HTTP client using Go.

```go
func init() {
    fmt.Println("Inside init function")
    if len(os.Args) != 2 {
        fmt.Println("Usage: gocurl <url>")
        os.Exit(1)
    }
}

func GoCurl() {
    fmt.Println("Inside GoCurl function")
    _, err := http.Get(os.Args[1])
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
```

**Note:** `init()` functions run automatically before `main()`, in the order of package imports.

## Message Passing

Channel-based message passing pattern.

```go
func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func MessagePassing() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs) // "passed message"
}
```

## Link Checker (Continuous Monitoring)

Continuously checks websites and reports status using goroutines and channels.

```go
func ChannelLinks() {
    links := []string{
        "http://google.com",
        "http://facebook.com",
        "http://golang.org",
        "http://amazon.com",
    }

    c := make(chan string, len(links))

    for _, link := range links {
        go checkLink(link, c)
    }

    // Continuously re-check links
    for l := range c {
        go func(link string) {
            time.Sleep(2 * time.Second)
            checkLink(link, c)
        }(l)
    }
}

func checkLink(link string, c chan string) {
    _, err := http.Get(link)
    if err != nil {
        fmt.Println(link, "might be down!")
    } else {
        fmt.Println(link, "is up!")
    }
    c <- link
}
```

**Key pattern:** A function literal (`func(link string)`) captures the loop variable explicitly to avoid the classic closure-in-a-loop trap.
