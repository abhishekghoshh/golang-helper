# Logging

Go provides the `log` package for free-form output and `log/slog` for structured output.

```go
import (
    "bytes"
    "log"
    "log/slog"
    "os"
)

func Loggings() {
    // Standard logger (writes to stderr)
    log.Println("standard logger")

    // Configure flags
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)
    log.Println("with micro")

    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Println("with file/line")

    // Custom logger
    customLog := log.New(os.Stdout, "custom-log:", log.LstdFlags)
    customLog.Println("from custom log")

    // Change prefix
    customLog.SetPrefix("this is prefix:")
    customLog.Println("with prefix")

    // Log to a buffer (any io.Writer works)
    var buf bytes.Buffer
    buflog := log.New(&buf, "buf:", log.LstdFlags)
    buflog.Println("hello")
    fmt.Print("from buflog:", buf.String())

    // Structured logging with slog (JSON output)
    jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
    myslog := slog.New(jsonHandler)
    myslog.Info("hi there")

    // With key-value pairs
    myslog.Info("hello again", "key", "val", "age", 25)
}
```

**Key points:**
- `log.Fatal*` and `log.Panic*` exit or panic after logging
- Default logger outputs to `os.Stderr` with date and time
- `slog` supports JSON output out of the box
- Custom loggers target any `io.Writer`
