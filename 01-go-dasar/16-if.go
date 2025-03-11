package main

import "fmt"

func main() {
	name := "Svelte"

	if name == "React" {
		fmt.Println("Hello if")
	} else if name == "Vue" {
		fmt.Println("Hello else if 1")
	} else if name == "Svelte" {
		fmt.Println("Hello else if 2")
	} else {
		fmt.Println("Hello else")
	}

	if length := len(name); length <= 5 {
		fmt.Println("panjanga name <= 5")
	} else {
		fmt.Println("panjanga name > 5")
	}
}