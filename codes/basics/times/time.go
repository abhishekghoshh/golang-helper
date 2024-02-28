package times

import (
	"fmt"
	"time"
)

/*
https://gobyexample.com/time

Go offers extensive support for times and durations; here are some examples.
*/
func Times() {
	print := fmt.Println

	//We'll start by getting the current time.
	now := time.Now()
	print("current time is", now)

	//You can build a time struct by providing the year, month, day, etc.
	// Times are always associated with a Location, i.e. time zone.
	birthday := time.Date(1997, 11, 17, 20, 34, 58, 651387237, time.UTC)
	print("constructed birthday is ", birthday)

	// You can extract the various components of the time value as expected.
	print(birthday.Year())
	print(birthday.Month())
	print(birthday.Day())
	print(birthday.Hour())
	print(birthday.Minute())
	print(birthday.Second())
	print(birthday.Nanosecond())
	print(birthday.Location())

	// The Monday-Sunday Weekday is also available.
	print(birthday.Weekday())

	// These methods compare two times, testing if the first occurs before, after,
	// or at the same time as the second, respectively.
	print(birthday.Before(now))
	print(birthday.After(now))
	print(birthday.Equal(now))

	// The Sub methods returns a Duration representing the interval between two times.
	diff := now.Sub(birthday)
	print("difference is ", diff)

	// We can compute the length of the duration in various units.
	print(diff.Hours())
	print(diff.Minutes())
	print(diff.Seconds())
	print(diff.Nanoseconds())

	// You can use Add to advance a time by a given duration, or with a - to move backwards by a duration.
	print(birthday.Add(diff))
	print(birthday.Add(-diff))
}
