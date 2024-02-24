package files

import (
	"bufio"
	"fmt"
	"os"
)

// https://gobyexample.com/writing-files
func WritingFiles() {

	var check = func(e error) {
		if e != nil {
			panic(e)
		}
	}

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("newData.txt", d1, 0644)
	check(err)

	f, err := os.Create("newData.txt")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
