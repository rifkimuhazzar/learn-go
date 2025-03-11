package main

import (
	"01-go-dasar/helper"
	"fmt"
) 

func main() {
	result := helper.SayHello("Vue")
	fmt.Println(result)

	fmt.Println(helper.Aplication)
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodBye())
	fmt.Println(helper.Example("Test"))
}