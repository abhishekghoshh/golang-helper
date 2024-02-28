package strings

import (
	"fmt"
	"strings"
)

/*
https://gobyexample.com/string-functions
https://pkg.go.dev/strings

The standard library's strings package provides many useful string-related functionstrings.
Here are some examples to give you a sense of the package.

Here’s a sample of the functions available in stringstrings. Since these are functions from the package,
not methods on the string object itself, we need pass the string in question as the first argument
to the function. You can find more functions in the strings package docstrings.
*/
func Stringfunctions() {
	fmt.Println("Contains:  ", strings.Contains("test", "es"))
	fmt.Println("Count:     ", strings.Count("test", "t"))
	fmt.Println("HasPrefix: ", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", strings.HasSuffix("test", "st"))
	fmt.Println("Index:     ", strings.Index("test", "e"))
	fmt.Println("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:    ", strings.Repeat("a", 5))
	fmt.Println("Replace:   ", strings.Replace("fooo", "o", "0", -1))
	fmt.Println("Replace:   ", strings.Replace("fooo", "o", "0", 1))
	fmt.Println("Replace:   ", strings.Replace("fooo", "o", "0", 2))
	fmt.Println("Split:     ", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   ", strings.ToLower("TEST"))
	fmt.Println("ToUpper:   ", strings.ToUpper("test"))
}
