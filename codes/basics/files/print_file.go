package files

import (
	"fmt"
	"io"
	"os"
)

func PrintFile() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
