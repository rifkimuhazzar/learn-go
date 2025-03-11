package main

import "fmt"

func main() {
	name := "Svelte"

	switch name {
		case "React":
			fmt.Println("Hello React")
		case "Vue":
			fmt.Println("Hello Vue")
		case "Svelte":
			fmt.Println("Hello Svelte")
		default:
			fmt.Println("Hello Unknown")
	}

	switch length := len(name); length <= 5 {
		case true:
			fmt.Println("panjang <= 5")
		case false:
			fmt.Println("panjang > 5")
	}

	
	switch length := len(name); {
		case length > 10:
			fmt.Println("panjang > 5")
		case length > 5:
			fmt.Println("panjang > 5")
		default:
			fmt.Println("panjang <= 5")
	}
}