package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	var s string
	for i := 0; i < 100000; i++ {
		s += "x"
	}
	fmt.Println("time elapsed", time.Since(startTime))

	startTime = time.Now()
	var buffer bytes.Buffer
	for i := 0; i < 100000; i++ {
		buffer.WriteString("x")
	}
	s = buffer.String()
	fmt.Println("time elapsed", time.Since(startTime))

	startTime = time.Now()
	var builder strings.Builder
	for i := 0; i < 100000; i++ {
		builder.WriteString("x")
	}
	s = builder.String()
	fmt.Println("time elapsed", time.Since(startTime))
}
