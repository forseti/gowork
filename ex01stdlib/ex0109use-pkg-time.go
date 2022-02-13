package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now =", now)
	fmt.Println("year =", now.Year())
	fmt.Println("month =", now.Month())
	fmt.Println("day =", now.Day())
	fmt.Println("hour =", now.Hour())
	fmt.Println("minute =", now.Minute())
	fmt.Println("second =", now.Second())
	fmt.Println("nanosecond =", now.Nanosecond())

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println("utc =", utc)

	parsed, _ := time.Parse(time.RFC3339, "2021-12-21T00:00:00Z")
	fmt.Println("parsed by time.RFC3339 =", parsed)
	parsed, _ = time.Parse(time.RFC3339, "2022-02-22T22:22:22+00:00")
	fmt.Println("parsed by time.RFC3339 =", parsed)
	layout := "2006-01-02"
	parsed, _ = time.Parse(layout, "2012-12-21")
	fmt.Println("parsed by \"2006-01-02\" =", parsed)
	layout = "02-01-2006 15:04:05"
	parsed, _ = time.Parse(layout, "21-12-2012 12:12:12")
	fmt.Println("parsed by \"02-01-2006 15:04:05\" =", parsed)
}
