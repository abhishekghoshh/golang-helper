package singleton

import (
	"design-pattern/src/creational/singleton/instance"
	"fmt"
)

func Main() {
	for i := 0; i < 30; i++ {
		go instance.GetInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
