package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked:", name)
	} else {
		fmt.Println("Weclome", name)
	}
}

func main() {
	blacklist := func (x string) bool {
		return x == "x"
	}
	registerUser("X", blacklist)
	
	registerUser("x", func (x string) bool {
		return x == "x"
	})
}