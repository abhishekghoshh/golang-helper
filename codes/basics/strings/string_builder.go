package strings

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

const STRING_LENGTH = 100000

func StringBuilers() {
	// there are three ways to create a string
	// 1. normal string
	// 2. Byte buffer
	// 3. String builder
	calculateTimeTaken(
		"normal string",
		func() {
			var s string
			for i := 0; i < STRING_LENGTH; i++ {
				s += "x"
			}
		},
	)

	calculateTimeTaken(
		"String with byte buffer",
		func() {
			var buffer bytes.Buffer
			for i := 0; i < STRING_LENGTH; i++ {
				buffer.WriteString("x")
			}
			_ = buffer.String()
		},
	)

	calculateTimeTaken(
		"String with string builder",
		func() {
			var builder strings.Builder
			for i := 0; i < STRING_LENGTH; i++ {
				builder.WriteString("x")
			}
			_ = builder.String()
		},
	)
}

func calculateTimeTaken(identifier string, fn func()) {
	startTime := time.Now()
	fn()
	fmt.Println(identifier, "time elapsed", time.Since(startTime))
}
