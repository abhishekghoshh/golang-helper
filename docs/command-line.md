# Command Line

## Command Line Arguments

`os.Args` provides raw command-line arguments. `os.Args[0]` is the program path; `os.Args[1:]` are the actual arguments.

```go
func CommandLineArguments() {
    argsWithProg := os.Args
    fmt.Println(argsWithProg)

    argsWithoutProg := os.Args[1:]
    fmt.Println(argsWithoutProg)

    arg := os.Args[3] // individual arg
    fmt.Println(arg)
}
```

Run with: `go run main.go arg1 arg2 arg3`

## Command Line Flags

The `flag` package supports string, integer, and boolean flags.

```go
func CommandLineFlags() {
    wordPtr := flag.String("word", "foo", "a string")
    numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")

    // Bind to existing variable
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args()) // positional args
}
```

Run with: `go run main.go -word=opt -numb=7 -fork -svar=flag tail1 tail2`

## Command Line Subcommands

Use `flag.NewFlagSet()` for subcommands like `git commit` or `go build`.

```go
func CommandLineSubCommands() {
    fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
    fooEnable := fooCmd.Bool("enable", false, "enable")
    fooName := fooCmd.String("name", "", "name")

    barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
    barLevel := barCmd.Int("level", 0, "level")

    if len(os.Args) < 2 {
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "foo":
        fooCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'foo'")
        fmt.Println("  enable:", *fooEnable)
        fmt.Println("  name:", *fooName)
        fmt.Println("  tail:", fooCmd.Args())
    case "bar":
        barCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'bar'")
        fmt.Println("  level:", *barLevel)
        fmt.Println("  tail:", barCmd.Args())
    default:
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }
}
```

Run with: `go run main.go foo -enable -name=joe a1 a2`

## Environment Variables

```go
func EnvironmentVariables() {
    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))  // "1"
    fmt.Println("BAR:", os.Getenv("BAR"))  // "" (not set)

    // List all environment variables
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0], pair[1])
    }
}
```

Run with env: `BAR="I am BAR" go run main.go`
