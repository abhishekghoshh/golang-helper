package main

import (
	"fmt"
	"strconv"
)

func greet(name string) string {
	return "Hello, " + name
}

// you can have more than 1 default set
func greetWithDefault(names ...string) string {
	if len(names) == 0 {
		return greet("Abhishek Ghosh")
	}
	name := names[0]
	return greet(name)
}

type GreetingOptions struct {
	Name string
	Age  int
}

func Greet(options GreetingOptions) string {
	return "Hello, my name is " + options.Name + " and I am " + strconv.Itoa(options.Age) + " years old."
}

// Functional options pattern starts
type GreetingOption func(*GreetingOptions)

func WithName(name string) GreetingOption {
	return func(o *GreetingOptions) {
		o.Name = name
	}
}

func WithAge(age int) GreetingOption {
	return func(o *GreetingOptions) {
		o.Age = age
	}
}

func GreetWithDefaultOptions(options ...GreetingOption) string {
	opts := GreetingOptions{
		Name: "Abhishek Ghosh",
		Age:  25,
	}
	for _, o := range options {
		o(&opts)
	}
	return Greet(opts)
} // Functional options pattern ends

func main() {
	fmt.Println(greetWithDefault())
	fmt.Println(greetWithDefault("Nasim Molla"))

	fmt.Println(GreetWithDefaultOptions(WithName("Alice"), WithAge(20)))

}
