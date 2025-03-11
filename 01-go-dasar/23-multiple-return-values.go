package main

import "fmt"

func getFullName() (string, string) {
	return "Hello", "World"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	first, _ := getFullName()
	fmt.Println(first)
}