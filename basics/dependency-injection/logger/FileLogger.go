package logger

import "os"

type FileLogger struct {
	File *os.File
}

func (l *FileLogger) Log(message string) error {
	_, err := l.File.Write([]byte(message))
	l.File.Write([]byte("\n"))
	return err
}
