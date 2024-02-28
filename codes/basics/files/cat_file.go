package files

import (
	"fmt"
	"io"
	"os"
)

func CatFile() {
	fmt.Println("I am here")
	file, err := fetchFilePath()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}

func fetchFilePath() (*os.File, error) {
	if len(os.Args) == 1 {
		var filePath string
		fmt.Scanln(&filePath)
		return os.Open(filePath)
	}
	return os.Open(os.Args[1])
}
