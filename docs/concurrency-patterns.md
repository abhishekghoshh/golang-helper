# Concurrency Patterns

## Worker Pool for File Processing

A practical example comparing single-threaded vs. concurrent file processing using a worker pool pattern.

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func watchFunction(identifier string, runnable func()) {
    startTime := time.Now()
    runnable()
    fmt.Println("total time elapsed for", identifier, time.Since(startTime))
}

// Single-threaded: reads and writes sequentially
func withoutWorkerPool() {
    fileSrc, _ := os.Open("customers.csv")
    defer fileSrc.Close()

    fileDest, _ := os.OpenFile("copy.csv", os.O_APPEND|os.O_WRONLY, 0600)
    defer fileDest.Close()

    scanner := bufio.NewScanner(fileSrc)
    for scanner.Scan() {
        data := scanner.Text()
        fileDest.WriteString(data + "\n")
    }
}

// Worker pool: reads into channel, workers write concurrently
func withWorkerPool() {
    capacity, workers := 100000, 1000
    ch := make(chan string, capacity)
    marker := make(chan bool)

    // Goroutine to signal when channel should close
    go func() {
        select {
        case <-marker:
            close(ch)
        default:
            continue
        }
    }()

    file, _ := os.Open("customers.csv")
    defer file.Close()

    // Start workers
    for i := 0; i < workers; i++ {
        go saveToFile(ch)
    }

    // Feed data into channel
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        ch <- scanner.Text()
    }
    marker <- true
}

func saveToFile(ch chan string) {
    f, _ := os.OpenFile("copy-workerpool.csv", os.O_APPEND|os.O_WRONLY, 0600)
    defer f.Close()
    for data := range ch {
        f.WriteString(data + "\n")
    }
}

func main() {
    watchFunction("without-worker-pool", withoutWorkerPool)
    watchFunction("with-worker-pool", withWorkerPool)
}
```

**Key pattern:** A buffered channel acts as a job queue. Multiple worker goroutines read from it and process concurrently. A marker channel signals when all jobs have been enqueued.
