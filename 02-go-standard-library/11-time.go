package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Local())

	utc := time.Date(2025, time.December, 17, 12, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2025-12-12 12:00:00"
	// value := "2020"
	valueTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println("if")
		fmt.Println(valueTime)
		fmt.Println(err)
	} else {
		fmt.Println("else")
		fmt.Println(valueTime)
		fmt.Println(valueTime.Local())
		fmt.Println(err)
	}
	
	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
	fmt.Println(valueTime.Minute())
	fmt.Println(valueTime.Second())
	fmt.Println(valueTime.Nanosecond())	
}