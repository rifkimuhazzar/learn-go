package main

import "fmt"

type Address struct {
	frontEnd, backEnd string
}

func main() {
	address := Address{"Svelte", "Go"}
	fmt.Println(address)

	address1 := Address{"React", "Express"}
	// address2 := address1 // pass by value - default
	address2 := &address1 // pass by reference - pointer
	address2.frontEnd = "Vue" // otomatis jadi (*address2).frontEnd = "Vue"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(*address2)

	fmt.Println("--------------------------")

	address3 := &address2
	fmt.Println(address3)
	fmt.Println(*address3)
	fmt.Println(**address3)

	fmt.Println("--------------------------")

	address4 := &address3
	fmt.Println(address4)
	fmt.Println(*address4)
	fmt.Println(**address4)
	fmt.Println(***address4)

	fmt.Println("--------------------------")
	
	address5 := &address
	fmt.Println(address5)
	fmt.Println(*address5)
	
	fmt.Println("--------------------------")

	fmt.Println(&address)
	fmt.Println(&address2)
	fmt.Println(&address3)
	fmt.Println(&address4)
	fmt.Println(&address5)
}