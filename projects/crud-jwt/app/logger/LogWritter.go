package logger

import (
	"fmt"
	"os"
)

type LogWritter struct {
	logPath string
}

func (logWritter *LogWritter) Write(p []byte) (n int, err error) {
	content := string(p)
	logWritter.appendToFile(logWritter.logPath, content)
	fmt.Println(content)
	return len(p), nil
}
func (logWritter *LogWritter) appendToFile(filename, data string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(data + "\n"); err != nil {
		panic(err)
	}
}
