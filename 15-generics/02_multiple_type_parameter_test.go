package generics

import (
	"fmt"
	"testing"
)

func multiple[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultiple(t *testing.T) {
	multiple[int, string](100, "Hello")
	multiple("Hello", 100)
}
