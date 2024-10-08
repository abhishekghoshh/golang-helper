package introduction

import (
	"fmt"
)

/*
https://gobyexample.com/defer

Defer is used to ensure that a function call is performed later in a program's execution,
usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.

Suppose we wanted to create a file, write to it, and then close when we're done.
Here's how we could do that with defer.
Immediately after getting a file object with createFile,
we defer the closing of that file with closeFile. This will be executed
at the end of the enclosing function (main), after writeFile has finished.

	func main() {
		f := createFile("/tmp/defer.txt")
		defer closeFile(f)
		writeFile(f)
	}
	func createFile(p string) *os.File {
		fmt.Println("creating")
		f, err := os.Create(p)
		if err != nil {
			panic(err)
		}
		return f
	}
	func writeFile(f *os.File) {
		fmt.Println("writing")
		fmt.Fprintln(f, "data")
	}
	func closeFile(f *os.File) {
		fmt.Println("closing")
		err := f.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}

It's important to check for errors when closing a file, even in a deferred function.
*/
func Defer() {
	defer run("main")
	fmt.Println("The main method starts")

	// anynomous function is called when it is declared
	func() {
		defer run("anynomous")
		fmt.Println("In the anonymous function")
	}()

	fmt.Println("counting started")

	for i := 0; i < 10; i++ {
		defer fmt.Print(i, " ")
	}

	fmt.Println("couting finished")

	// last statement of main method
	fmt.Println("The main method finished")

}

func run(context string) {
	fmt.Println("\t deffered run method is executed inside", context)
}
