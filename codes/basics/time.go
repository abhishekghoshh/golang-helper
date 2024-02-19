package main

import (
	"fmt"
	"time"
)

// https://gobyexample.com/time
func main() {
	print := fmt.Println

	now := time.Now()
	print(now)

	birthday := time.Date(1997, 11, 17, 20, 34, 58, 651387237, time.UTC)
	print(birthday)

	print(birthday.Year())
	print(birthday.Month())
	print(birthday.Day())
	print(birthday.Hour())
	print(birthday.Minute())
	print(birthday.Second())
	print(birthday.Nanosecond())
	print(birthday.Location())

	print(birthday.Weekday())

	print(birthday.Before(now))
	print(birthday.After(now))
	print(birthday.Equal(now))

	diff := now.Sub(birthday)
	print(diff)

	print(diff.Hours())
	print(diff.Minutes())
	print(diff.Seconds())
	print(diff.Nanoseconds())

	print(birthday.Add(diff))
	print(birthday.Add(-diff))
}
