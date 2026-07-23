# Getting Started with Go

Go is an open-source programming language designed at Google in 2007 to improve programming productivity in an era of multicore, networked machines and large codebases. The language is often referred to as Golang because of its domain name, golang.org, but its proper name is Go.

## Hello World

```go
package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

var deckSize int

func main() {
    deckSize = 50
    fmt.Println(deckSize)

    fmt.Println("Hello World")
    fmt.Println("The time is", time.Now())
    fmt.Println("My favourite number is", rand.Intn(100))
    fmt.Printf("Sqrt of 7 is %g \n", math.Sqrt(7))
    fmt.Println("value of PI is", math.Pi)
}
```

Package `main` is the entry point for a Go program. The `main()` function is where execution begins.

## Module Setup

```bash
# Enable Go modules
export GO111MODULE=on

# Initialize a new module
go mod init basics

# Add dependencies
go get github.com/gorilla/mux
```

A `go.mod` file at the module root defines the module path and its dependencies. Example:

```
module basics

go 1.22.0

require (
    github.com/gorilla/mux v1.8.1
    golang.org/x/tour v0.1.0
)
```

## Go Commands

- **go build** — Compiles source code files into a binary
- **go run** — Compiles and executes one or more files
- **go fmt** — Formats all code in each file in the current directory
- **go install** — Compiles and installs a package in the bin directory
- **go get** — Downloads the raw source code of someone else's package
- **go test** — Runs any tests associated with the current project
