package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)
	// data.Value = "1"

	// data = data.Next()
	// data.Value = "2"

	// data = data.Next()
	// data.Value = "3"

	// data = data.Next()
	// data.Value = "4"

	// data = data.Next()
	// data.Value = "5"

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i + 1)
		data = data.Next()
	}
	
	fmt.Println(data)
	fmt.Println(data.Next())
	fmt.Println(data.Next().Next().Next().Next().Next())

	data.Do(func (value any) {
		fmt.Println(value)
	})
}