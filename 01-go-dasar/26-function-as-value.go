package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	fmt.Println(getGoodBye("React"))
	
	goodbye := getGoodBye
	fmt.Println(goodbye("Vue"))

	fmt.Println(getGoodBye("Svelte"))
}