package introduction

import (
	"fmt"
	"math"
)

/*
An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods.

A type implements an interface by implementing its methods.
There is no explicit declaration of intent, no "implements" keyword.
Implicit interfaces decouple the definition of an interface from its implementation,
which could then appear in any package without prearrangement.


Under the hood, interface values can be thought of as a tuple of a value and a concrete type: (value, type)
An interface value holds a value of a specific underlying concrete type.
Calling a method on an interface value executes the method of the same name on its underlying type.


If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
In some languages this would trigger a null pointer exception,
but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)
Note that an interface value that holds a nil concrete value is itself non-nil.


A nil interface value holds neither value nor concrete type.
Calling a method on a nil interface is a run-time error
because there is no type inside the interface tuple to indicate which concrete method to call.

	type I interface {
		M()
	}

	func main() {
		var i I
		describe(i)
		i.M()
	}

	func describe(i I) {
		fmt.Printf("(%v, %T)\n", i, i)
	}

The interface type that specifies zero methods is known as the empty interface: interface{}
An empty interface may hold values of any type. (Every type implements at least zero methods.)
Empty interfaces are used by code that handles values of unknown type.
For example, fmt.Print takes any number of arguments of type interface{}.

A type assertion provides access to an interface value's underlying concrete value.t := i.(T)
This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
If i does not hold a T, the statement will trigger a panic.
To test whether an interface value holds a specific type, a type assertion can return two values:
the underlying value and a boolean value that reports whether the assertion succeeded. t, ok := i.(T)
If i holds a T, then t will be the underlying value and ok will be true.
If not, ok will be false and t will be the zero value of type T, and no panic occurs.
Note the similarity between this syntax and that of reading from a map.
*/

// https://gobyexample.com/interfaces
func Interfaces() {
	r := rect{
		width:  3,
		height: 4,
	}
	c := circle{
		radius: 5,
	}

	measure(r)
	measure(c)

	bot1 := ChatGPT{}
	write(bot1)

	bot2 := Bard{}
	getReply(bot2)

	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	// t is non nil but the value inside it is nil
	var t *T
	i = t
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	// type assertions

	var empty_i interface{} = "hello"

	s := empty_i.(string)
	fmt.Println(s)

	s, ok := empty_i.(string)
	fmt.Println(s, ok)

	f, ok := empty_i.(float64)
	fmt.Println(f, ok)

	do(21)
	do("hello")
	do(true)
	do(nil)
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}
type geometry interface {
	area() float64
	perim() float64
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

type Bot interface {
	getReply() string
}

func getReply(bot Bot) {
	fmt.Println(bot.getReply())
}

type Writter interface {
	write()
}
type BotWritter interface {
	Bot
	Writter
}

func write(botWritter BotWritter) {
	botWritter.write()
}

type ChatGPT struct{}
type Bard struct{}

func (ChatGPT) getReply() string {
	return "Hi I'am ChatGPT"
}
func (bot ChatGPT) write() {
	fmt.Println("Hi I am bot Writter and you are working with ChatGPT")
}

func (Bard) getReply() string {
	return "Hi I'm Bard"
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
