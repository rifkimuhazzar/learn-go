package main

import "fmt"

// func sayHelloWithFilter(name string, filter func(param1 string) string) {
// 	filteredName := filter(name)
// 	fmt.Println("Hello", filteredName)
// }

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "x" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("X", spamFilter)

	test := spamFilter
	sayHelloWithFilter("x", test) 
}
