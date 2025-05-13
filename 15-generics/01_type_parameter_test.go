package generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func printValue[x any](value x) x {
	fmt.Println(value)
	return value
}

func TestPrintValue(t *testing.T) {
	result_1 := printValue("Svelte")
	assert.Equal(t, "Svelte", result_1)

	result_2 := printValue[int](100)
	assert.Equal(t, 100, result_2)

	result_3 := printValue[any](100)
	assert.Equal(t, 100, result_3)
}
