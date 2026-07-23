# Time and Date

## Current Time

```go
func Times() {
    now := time.Now()
    fmt.Println("current time is", now)

    // Construct a specific time
    birthday := time.Date(1997, 11, 17, 20, 34, 58, 651387237, time.UTC)

    // Extract components
    fmt.Println(birthday.Year())       // 1997
    fmt.Println(birthday.Month())      // November
    fmt.Println(birthday.Day())        // 17
    fmt.Println(birthday.Hour())       // 20
    fmt.Println(birthday.Minute())     // 34
    fmt.Println(birthday.Second())     // 58
    fmt.Println(birthday.Weekday())    // Monday

    // Comparisons
    fmt.Println(birthday.Before(now))  // true
    fmt.Println(birthday.After(now))   // false

    // Duration between times
    diff := now.Sub(birthday)
    fmt.Println(diff.Hours())
    fmt.Println(diff.Minutes())

    // Add / subtract durations
    fmt.Println(birthday.Add(diff))
    fmt.Println(birthday.Add(-diff))
}
```

## Epoch (Unix Time)

```go
func Epochs() {
    now := time.Now()

    fmt.Println(now.Unix())       // seconds since epoch
    fmt.Println(now.UnixMilli())  // milliseconds
    fmt.Println(now.UnixNano())   // nanoseconds

    // Convert back from epoch
    fmt.Println(time.Unix(now.Unix(), 0))
    fmt.Println(time.Unix(0, now.UnixNano()))
}
```

## Time Formatting and Parsing

Go uses a reference time for format layouts: `Mon Jan 2 15:04:05 MST 2006` (Unix epoch date `01/02 03:04:05PM '06 -0700`).

```go
func TimeFormatting() {
    t := time.Now()

    // Standard formats
    fmt.Println(t.Format(time.RFC3339))          // 2024-01-15T10:30:00Z

    // Custom layouts (use the reference time)
    fmt.Println(t.Format("3:04PM"))               // 10:30AM
    fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
    fmt.Println(t.Format("2006-01-02T15:04:05.999999-07:00"))

    // Parsing
    t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
    fmt.Println(t1)

    form := "3 04 PM"
    t2, _ := time.Parse(form, "8 41 PM")
    fmt.Println(t2)

    // Manual formatting with Printf
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())
}
```

**Reference time components:** `01/02 03:04:05PM '06 -0700` maps to `Jan 2 3:04:05PM 2006 MST`.
