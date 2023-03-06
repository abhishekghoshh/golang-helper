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
}
