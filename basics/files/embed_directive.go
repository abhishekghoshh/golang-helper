package files

import (
	"embed"
)

/*
https://gobyexample.com/embed-directive
https://pkg.go.dev/embed
https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives
https://pkg.go.dev/embed#FS

//go:embed is a compiler directive that allows programs to include arbitrary
files and folders in the Go binary at build time. Read more about the embed directive here.

Import the embed package; if you don't use any exported identifiers from this package, you can do a blank import with _ "embed".

*/

// embed directives accept paths relative to the directory containing the Go source file.
// This directive embeds the contents of the file into the string variable immediately following it.

//go:embed embed/single_file.txt
var fileString string

// We can also embed multiple files or even folders with wildcards.
// This uses a variable of the embed.FS type, which implements a simple virtual file system.
// Or embed the contents of the file into a []byte.

//go:embed embed/single_file.txt
var fileByte []byte

//go:embed embed/single_file.txt
//go:embed embed/*.hash
var folder embed.FS

func EmbeddedDirective() {

	// Print out the contents of single_file.txt.
	print(fileString)
	println("embed file content", string(fileByte))

	// Retrieve some files from the embedded folder.
	content1, _ := folder.ReadFile("embed/file1.hash")
	println("read files content for file 1 :", string(content1))

	content2, _ := folder.ReadFile("embed/file2.hash")
	println("read files content for file 2 :", string(content2))

	/*

		$ mkdir -p ./files/embed
		$ echo "random file content" > ./files/embed/single_file.txt
		$ echo "123" > ./files/embed/file1.hash
		$ echo "456" > ./files/embed/file2.hash
		$ go run main.go

	*/

}
