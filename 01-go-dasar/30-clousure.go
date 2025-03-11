package main

import "fmt"

func main() {
	counter := 0
	increment := func ()  {
		fmt.Println("Increment 1")
		counter++
	}

	increment()
	increment()
	increment()
	increment()
	increment()
	fmt.Println(counter)
}