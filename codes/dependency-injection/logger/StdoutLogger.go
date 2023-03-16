package logger

import "fmt"

type StdoutLogger struct{}

func (l *StdoutLogger) Log(message string) error {
	_, err := fmt.Println(message)
	return err
}
