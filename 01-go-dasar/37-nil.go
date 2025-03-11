package main

import "fmt"

// nill only support for interface, function, map, slice, pointer, channel
// func Example(name string) string {
// 	return nil
// }

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("React")
	fmt.Println(data)
	fmt.Println(data["name"])
	if data == nil {
		fmt.Println("data is empty")
	}
}