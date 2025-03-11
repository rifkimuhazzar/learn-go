package main

import "fmt"

type Address struct {
	frontEnd, backEnd string
}

func main() {
	// address1 := &Address{}
	address1 := new(Address)
	address2 := address1

	fmt.Println(address1)
	fmt.Println(address2)

	address2.backEnd = "Express"

	fmt.Println(address1)
	fmt.Println(address2)

	fmt.Println("----------------------------")

	ptr := new(int) 
	fmt.Println(ptr) 
	fmt.Println(*ptr) 

	*ptr = 10 
	fmt.Println(ptr) 
	fmt.Println(*ptr) 

}
