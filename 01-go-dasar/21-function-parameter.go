package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, "and", lastName)
}

func main() {
	sayHelloTo("React", "Vue")
	sayHelloTo("Next", "Nuxt")
}