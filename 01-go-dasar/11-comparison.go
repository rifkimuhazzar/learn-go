package main

import "fmt"

func main() {
	name1 := "Hello Go1"
	name2 := "Hello Go"

	var result1 = name1 == name2
	var result2 = name1 != name2
	var result3 = name1 > name2

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}