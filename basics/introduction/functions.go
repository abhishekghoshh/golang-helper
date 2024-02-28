package introduction

import (
	"fmt"
	"strconv"
)

/*
Go requires explicit returns, i.e. it won't automatically return the value of the last expression.

When you have multiple consecutive parameters of the same type,
you may omit the type name for the like-typed parameters up to the final parameter that declares the type.


Go has built-in support for multiple return values. This feature is used often in idiomatic Go,
for example to return both result and error values from a function.

The (int, int) in this function signature shows that the function returns 2 ints.
Here we use the 2 different return values from the call with multiple assignment.
If you only want a subset of the returned values, use the blank identifier _
*/

func Functions() {

	fmt.Println("addition of 1 and 3 is", add(1, 3))

	x, y := 5, 6
	fmt.Println("current value of x is", x, "current calue of y is", y)
	x, y = swap(x, y)
	fmt.Println("After swappingcurrent value of x is", x, "current calue of y is", y)

	welcome("Abhishek Ghosh")

	ageInDays := age("10", "09", "1997")
	fmt.Println("your age in days", ageInDays)

	abhishek := Human{
		Name: &Name{
			FirstName: "Abhishek",
			LastName:  "Ghosh",
		},
		Address: &Address{
			StreetName: "18 no alep khan Mahalla road",
			City:       "Berhampore",
			District:   "Murshidabad",
			State:      "West Bengal",
			Country:    "India",
			Pin:        742101,
		},
	}

	printPerson(&abhishek)

	print(abhishek)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func welcome(name string) {
	fmt.Println("Welcome", name)
}

func age(days, months, years string) int {
	days_, _ := strconv.Atoi(days)
	months_, _ := strconv.Atoi(months)
	years_, _ := strconv.Atoi(years)
	return days_ + (months_-1)*30 + years_*365
}

type Human struct {
	Name    *Name
	Address *Address
}
type Name struct {
	FirstName string
	LastName  string
}
type Address struct {
	StreetName string
	City       string
	District   string
	State      string
	Country    string
	Pin        int
}

func printPerson(person *Human) {
	name := person.Name
	address := person.Address

	fmt.Println("name of the person is", name.FirstName, name.LastName)
	fmt.Println("address of the person is", address.StreetName, address.City, address.District, address.State, address.Country, address.Pin)
}

func print(data any) {
	fmt.Println("printing data when type is not defined", data)
}
