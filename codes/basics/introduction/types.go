package introduction

import (
	"fmt"
	"math"
)

/*

we can set a particular type to our own custom type
and also attach some functionalities to them


You can declare a method on non-struct types, too.
In this example we see a numeric type MyFloat with an Abs method.
You can only declare a method with a receiver whose type is defined in the same package as the method.
You cannot declare a method with a receiver whose type is defined in another package

	type MyFloat float64

	func (f MyFloat) Abs() float64 {
		if f < 0 {
			return float64(-f)
		}
		return float64(f)
	}
	func main() {
		f := MyFloat(-math.Sqrt2)
		fmt.Println(f.Abs())
	}
*/

func Types() {
	c := color("Red")
	fmt.Println(c.describe("is an awesome color"))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
