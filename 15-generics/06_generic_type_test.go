package generics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func printBag[Z any](bag Bag[Z]) {
	for i, v := range bag {
		fmt.Println(i, ":", v)
	}
}

func TestPrintBag(t *testing.T) {
	numbersBag := Bag[int8]{10, 20, 30, 40, 50}
	fmt.Println(numbersBag)
	printBag(numbersBag)

	stringsBag := Bag[string]{"React", "Angular", "Vue", "Svelte", "Astro"}
	fmt.Println(stringsBag)
	printBag(stringsBag)
}
