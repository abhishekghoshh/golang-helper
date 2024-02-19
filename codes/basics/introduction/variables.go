package introduction

import (
	"fmt"
	"math/cmplx"
)

/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32

	// represents a Unicode code point

float32 float64

complex64 complex128

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.


Constants are declared like variables, but with the const keyword.
Constants cannot be declared using the := syntax.
*/

const Pi = 3.14

func Variables() {

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// zero value of the variables
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("Printing the zero values of these variables %v %v %v %q\n", i, f, b, s)

	fmt.Println("PI is constant", Pi)
}
