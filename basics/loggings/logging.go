package loggings

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

/*
https://gobyexample.com/logging

The Go standard library provides straightforward tools for outputting logs from Go programs,
with the log package for free-form output and the log/slog package for structured output.
*/
func Loggings() {

	// Simply invoking functions like Println from the log package uses the standard logger,
	// which is already pre-configured for reasonable logging output to os.Stderr.
	// Additional methods like Fatal* or Panic* will exit the program after logging.
	log.Println("standard logger")

	// Loggers can be configured with flags to set their output format. By default,
	// the standard logger has the log.Ldate and log.Ltime flags set, and these are collected in
	// log.LstdFlags. We can change its flags to emit time with microsecond accuracy, for example.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// It also supports emitting the file name and line from which the log function is called.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// It may be useful to create a custom logger and pass it around.
	// When creating a new logger, we can set a prefix to distinguish its output from other loggers.
	customLog := log.New(os.Stdout, "custom-log:", log.LstdFlags)
	customLog.Println("from custom log")

	// We can set the prefix on existing loggers (including the standard one) with the SetPrefix method.
	customLog.SetPrefix("this is prefix:")
	customLog.Println("This log is with prefix")

	// Loggers can have custom output targets; any io.Writer works.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// This call writes the log output into buf.
	buflog.Println("hello")

	// This will actually show it on standard output.
	fmt.Print("from buflog:", buf.String())

	// The slog package provides structured log output. For example, logging in JSON format is straightforward.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// In addition to the message, slog output can contain an arbitrary number of key=value pairs
	myslog.Info("hello again", "key", "val", "age", 25)
}
