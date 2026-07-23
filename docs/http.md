# HTTP Client & Server

## HTTP Client

```go
func HTTPClient() {
    resp, err := http.Get("https://gobyexample.com")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    // Print first 5 lines
    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }
}
```

## HTTP Server

```go
func HTTPServer() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "welcome to go server\n")
    })
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    fmt.Println("starting go server on 8090 port")
    http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}
```

## Context (Cancellation)

```go
func helloRoute(w http.ResponseWriter, req *http.Request) {
    ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

    select {
    case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "hello\n")
    case <-ctx.Done():
        err := ctx.Err()
        fmt.Println("server:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/hello", helloRoute)
    http.ListenAndServe(":8090", nil)
}
```

The context is cancelled when the client disconnects. Hit `Ctrl+C` on the client to see cancellation.

## Custom HTTP Response Writer

Implement the `io.Writer` interface to process HTTP response data.

```go
type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
    fmt.Println(string(bs))
    fmt.Println("wrote this many bytes:", len(bs))
    return len(bs), nil
}

func CustomHTTPWritter() {
    resp, _ := http.Get("http://google.com")

    // Read into byte slice
    byteSlice := make([]byte, 99999)
    n, _ := resp.Body.Read(byteSlice)
    fmt.Print(string(byteSlice[:n]))

    // Copy to os.Stdout using io.Copy
    io.Copy(os.Stdout, resp.Body)

    // Copy to custom writer
    lw := logWriter{}
    io.Copy(lw, resp.Body)
}
```
