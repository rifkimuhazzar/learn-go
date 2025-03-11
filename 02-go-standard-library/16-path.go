package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("hello/test/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello", "test", "world.go", "A"))
}