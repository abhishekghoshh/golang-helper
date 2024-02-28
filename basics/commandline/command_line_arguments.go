package commandline

import (
	"fmt"
	"os"
)

/*
https://gobyexample.com/command-line-arguments
https://en.wikipedia.org/wiki/Command-line_interface#Arguments

Command-line arguments are a common way to parameterize execution of programs.
For example, go run hello.go uses run and hello.go arguments to the go program.
*/
func CommandLineArguments() {

	// os.Args provides access to raw command-line arguments.
	// Note that the first value in this slice is the path to the program,
	// and os.Args[1:] holds the arguments to the program.
	argsWithProg := os.Args
	fmt.Println(argsWithProg)

	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	// You can get individual args with normal indexing.
	arg := os.Args[3]
	fmt.Println(arg)

	// To experiment with command-line arguments it's best to build a binary with go build first.
}
