package di

import (
	"fmt"
	"net/http"
	"os"
)

type Logger interface {
	Log(message string) error
}

type FileLogger struct {
	File *os.File
}

func (l *FileLogger) Log(message string) error {
	_, err := l.File.Write([]byte(message))
	l.File.Write([]byte("\n"))
	return err
}

type StdoutLogger struct{}

func (l *StdoutLogger) Log(message string) error {
	_, err := fmt.Println(message)
	return err
}

type Server struct {
	Logger Logger
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Logger.Log("received request from " + r.Host)
	s.Logger.Log("requested path is " + r.URL.Path)
}

func DependencyInjection() {
	file, _ := os.OpenFile("di/app.log", os.O_APPEND|os.O_WRONLY, 0600)
	fileLogger := &FileLogger{File: file}

	server := &Server{Logger: fileLogger}
	// server := &Server{Logger: &StdoutLogger{}}

	http.ListenAndServe(":8080", server)
}
