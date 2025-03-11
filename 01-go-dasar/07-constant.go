package main

import "fmt"

func main() {
	// const firstName string = "Hello"
	// const lastName = "World"

	const (
		firstName string = "Hello"
		lastName = "World"
	)

	// cannot reassign const
	// firstName = "Test"
	// lastName = "Test"

	fmt.Println(firstName, lastName)
}