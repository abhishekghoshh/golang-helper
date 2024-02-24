package files

import (
	"embed"
)

// go:embed folder/single_file.txt
var fileString string

// go:embed folder/single_file.txt
var fileByte []byte

// go:embed folder/single_file.txt
// go:embed folder/*.hash
var folder embed.FS

// https://gobyexample.com/embed-directive
func EmbeddedDirective() {

	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}

/*

$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash
$ go run embed-directive.go

*/
