package main

import "fmt"

func random() interface{} {
	return "Hello World"
}

func main() {
	var result = random()
	// var resultString = result.(string)
	// fmt.Println(result)
	// fmt.Println(resultString)

	// var resultInt = result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("string:", value)
	case int:
		fmt.Println("int:", value)
	default:
		fmt.Println("unknown:", value)
	}
}