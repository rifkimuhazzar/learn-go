package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "First"
	middleName = "Middle"
	// lastName = "Last"
	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}