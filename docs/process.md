# Process Management

## Spawning Processes

Run external commands and capture output.

```go
func DoSpawningProcess() {
    // Simple command
    dateCmd := exec.Command("date")
    dateOut, _ := dateCmd.Output()
    fmt.Println(string(dateOut))

    // Handle errors (wrong path / non-zero exit)
    _, err := exec.Command("date", "-x").Output()
    if err != nil {
        switch e := err.(type) {
        case *exec.Error:      fmt.Println("failed executing:", err)
        case *exec.ExitError:  fmt.Println("exit code =", e.ExitCode())
        }
    }

    // Pipe stdin and capture stdout
    grepCmd := exec.Command("grep", "hello")
    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := io.ReadAll(grepOut)
    grepCmd.Wait()
    fmt.Println(string(grepBytes)) // "hello grep"

    // Full command string with bash -c
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, _ := lsCmd.Output()
    fmt.Println(string(lsOut))
}
```

## Execing Processes (Replace current process)

Replace the current Go process entirely — execution stops after success.

```go
func DoExecingProcess() {
    binary, _ := exec.LookPath("ls")
    args := []string{"ls", "-a", "-l", "-h"}
    env := os.Environ()

    // This replaces the current process — nothing after will run
    syscall.Exec(binary, args, env)
}
```

Note: Go does not offer a classic Unix `fork`. Goroutines + spawning/execing covers most fork use cases.

## Signals

Handle Unix signals (SIGINT, SIGTERM) for graceful shutdown.

```go
func DoSignals() {
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    done := make(chan bool, 1)
    go func() {
        sig := <-sigs
        fmt.Println("signal:", sig)
        done <- true
    }()

    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
```

## Exit

`os.Exit` immediately exits with a status code. **Deferred functions do NOT run.**

```go
func DoExit() {
    defer fmt.Println("!") // never printed
    os.Exit(3)
}
```
