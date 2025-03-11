package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"React", "Vue", "Angular"}
	values := []int{100, 90, 80, 70, 100}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Vue"))
	fmt.Println(slices.Index(names, "vue"))
	fmt.Println(slices.Index(names, "Angular"))
}