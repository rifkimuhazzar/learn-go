package main

import "fmt"

// func test() interface{} {
// 	return 1
// 	return "test"
// 	return true
// }

func test() any {
	// return 1
	// return "test"
	return true
}

func main() {
	test := test()
	fmt.Println(test)
}