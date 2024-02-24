package process

import (
	"os"
	"os/exec"
	"syscall"
)

// https://gobyexample.com/execing-processes
func DoExecingProcess() {

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
