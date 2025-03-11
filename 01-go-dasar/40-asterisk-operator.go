package main

import "fmt"

type Address struct {
	frontEnd, backEnd string
}

func main() {
	address1 := Address{"React", "Express"}
	address2 := &address1
	address2.frontEnd = "Vue"
	fmt.Println(address1)
	fmt.Println(address2)

	// address2 = &Address{"Svelte", "Go"}
	*address2 = Address{"Svelte", "Go"}
	fmt.Println(address1)
	fmt.Println(address2)

	fmt.Println("-------------------------------")

	name1 := "React"
	name2 := &name1
	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(*name2)

	hello := "Hello"
	// name2 = &hello
	*name2 = hello
	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(*name2)

	a := "A"
	a = "B"
	a = "C"
	fmt.Println(a)
}