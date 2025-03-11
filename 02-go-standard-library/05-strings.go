package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello World", " Wo"))
	fmt.Println(strings.Split("Hello World", " "))
	fmt.Println(strings.ToLower("Hello World"))
	fmt.Println(strings.ToUpper("Hello World"))
	fmt.Println(strings.Trim("   Hello World     ", " "))
	fmt.Println(strings.ReplaceAll("Hello World, Hello Go", "Hello", "Test"))
}