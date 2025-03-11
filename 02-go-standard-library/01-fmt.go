package main

import "fmt"

func main() {
	first_name := "Hello"
	last_name := "World"
	age := 25

	fmt.Println(first_name, last_name, "'", age, "'")
	fmt.Printf("%s %s '%d'\n", first_name, last_name, age)
	fmt.Println(first_name, last_name, "'", age, "'")
}