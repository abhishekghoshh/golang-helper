package files

import (
	"fmt"
	"os"
	"time"
)

func AppendingFiles() {
	fmt.Println(appendToFile("./files/data/data.txt", time.Now().String()))
}
func appendToFile(filename, newData string) string {

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(newData + "\n"); err != nil {
		panic(err)
	}

	data, err_ := os.ReadFile(filename)
	if err_ != nil {
		fmt.Println(err)
	}
	return string(data)
}
