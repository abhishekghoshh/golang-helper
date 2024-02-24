package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://gobyexample.com/line-filters
func LineFilters() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// cat data.json | go run line-filters.go
