package main

import (
	"di/logger"
	"di/server"
	"net/http"
	"os"
)

func main() {
	file, _ := os.OpenFile("App.log", os.O_APPEND|os.O_WRONLY, 0600)
	logger := &logger.FileLogger{File: file}
	server := &server.Server{Logger: logger}
	http.ListenAndServe(":8080", server)
}
