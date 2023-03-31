package main

import (
	"fmt"
	"strconv"
)

func age(days, months, years string) int {
	days_, _ := strconv.Atoi(days)
	months_, _ := strconv.Atoi(months)
	years_, _ := strconv.Atoi(years)
	return days_ + (months_-1)*30 + years_*365
}

func welcome(name string) {
	fmt.Println("Welcome", name)
}
func main() {
	welcome("Abhishek Ghosh")
	ageInDays := age("10", "09", "1997")
	fmt.Println("your age in days", ageInDays)

	abhishek := Person{
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

	fmt.Println(abhishek)
	print(abhishek)
}

type Person struct {
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

func print(data any) {
	fmt.Println(data)
}
