package introduction

import (
	"fmt"
)

func run(context string) {
	fmt.Println("\t deffered run method is executed inside", context)
}

// https://gobyexample.com/defer
func Defer() {
	defer run("main")
	fmt.Println("The main method starts")

	// anynomous function is called when it is declared
	func() {
		defer run("anynomous")
		fmt.Println("In the anonymous function")
	}()

	fmt.Println("counting started")

	for i := 0; i < 10; i++ {
		defer fmt.Print(i, " ")
	}

	fmt.Println("couting finished")

	// last statement of main method
	fmt.Println("The main method finished")

}
