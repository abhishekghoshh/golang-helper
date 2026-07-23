# Testing in Go

## Steps for Writing Test Suites in Golang

- Create a file whose name ends with `_test.go`
- Import package testing by import `"testing"` command
- Write the test function of form `func TestXxx(*testing.T)` which uses any of Error, Fail, or related methods to signal failure.
- Put the file in any package.
- Run command `go test`
- Create go module
- Exported method names should be starting with capital case
- We should write all the clean-up codes in `t.Cleanup` method

**Note**: test file will be excluded in package build and will only get executed on `go test` command.

## Example

```go
import (
    "fmt"
    "testing"
)

func TestReturnGeeks(t *testing.T) {
    actualString := SayHello()
    expectedString := "hello"
    if actualString != expectedString {
        t.Errorf("Expected String(%s) is not same as"+
            " actual string (%s)", expectedString, actualString)
    }
    t.Cleanup(func() { fmt.Println("This is cleanup function") })
}
```

```go
import "fmt"

func SayHello() string {
    return "hello"
}

// main function of package
func main() {
    fmt.Println(SayHello())
}
```
