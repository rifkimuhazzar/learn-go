package main

import "fmt"

func main() {
	type NoKTP string

	var first NoKTP = "01010101"
	fmt.Println(first)

	var second string = "02020202"
	fmt.Println(second)

	var secondToNoKTP = NoKTP(second)
	fmt.Println(secondToNoKTP)
}