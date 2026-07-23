# Dependency Injection

Manual dependency injection using interfaces in Go. Inject dependencies via struct fields for testability and loose coupling.

```go
type Logger interface {
    Log(message string) error
}

// File-based logger
type FileLogger struct {
    File *os.File
}

func (l *FileLogger) Log(message string) error {
    _, err := l.File.Write([]byte(message + "\n"))
    return err
}

// Stdout logger
type StdoutLogger struct{}

func (l *StdoutLogger) Log(message string) error {
    _, err := fmt.Println(message)
    return err
}

// Server depends on Logger interface, not concrete type
type Server struct {
    Logger Logger
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s.Logger.Log("received request from " + r.Host)
    s.Logger.Log("requested path is " + r.URL.Path)
}

func main() {
    file, _ := os.OpenFile("app.log", os.O_APPEND|os.O_WRONLY, 0600)
    fileLogger := &FileLogger{File: file}

    server := &Server{Logger: fileLogger}
    // server := &Server{Logger: &StdoutLogger{}} // swap implementations

    http.ListenAndServe(":8080", server)
}
```

## Testing with DI

The `Logger` interface makes testing trivial — inject a mock.

```go
type MockLogger struct {
    Logs []string
}

func (l *MockLogger) Log(message string) error {
    l.Logs = append(l.Logs, message)
    return nil
}

func Test_DependencyInjection(t *testing.T) {
    logger := &MockLogger{}
    server := &Server{Logger: logger}

    // Trigger a request...
    // Assert on logger.Logs...
}
```

**Key principle:** Depend on interfaces, not concrete types. Swap implementations without changing business logic.
