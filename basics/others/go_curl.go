package others

import (
	"fmt"
	"net/http"
	"os"
)

// it is an example how we can write a command line application using go application
func GoCurl() {
	fmt.Println("Inside GoCurl function")
	_, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

// Whenever the init() function is declared in code,
// Go loads and runs it prior to anything else in the package
func init() {
	fmt.Println("Inside init function")
	if len(os.Args) != 2 {
		fmt.Println("Usage: gocurl <url>")
		os.Exit(1)
	}
}
