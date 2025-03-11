package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Hello React"
	names[1] = "Vue"
	names[2] = "Svelte"
	names[1] = "Hello Vue"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	values := [...]int{ 
		10, 
		20, 
		30,
		40,
		50,
	}
	values[1] = 25

	fmt.Println(values)
	fmt.Println(len(values))
}