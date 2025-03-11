package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("for 1:", i)
	}

	fmt.Println("--------------------------------")

	for i := 0; i < 10; i++ {
		if(i % 2 == 0) {
			continue
		}
		fmt.Println("for 2:", i)
	}
}