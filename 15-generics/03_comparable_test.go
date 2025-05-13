package generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func comp[T comparable](param1, param2 T) bool {
	if param1 == param2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {
	result1 := comp("Hello", "Hello")
	fmt.Println(result1)

	result2 := comp("Hello", "hello")
	fmt.Println(result2)

	assert.Equal(t, true, comp[string]("Svelte", "Svelte"))
	assert.Equal(t, true, comp(100, 100))
}
