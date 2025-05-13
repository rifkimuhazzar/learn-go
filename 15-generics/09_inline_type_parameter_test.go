package generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findMin[T interface{ int | int64 | float64 }](value1 T, value2 T) T {
	if value1 < value2 {
		return value1
	} else {
		return value2
	}
}

func TestInlineInterface(t *testing.T) {
	fmt.Println(findMin(100, 200))
	fmt.Println(findMin(100, 200.0))
	fmt.Println(findMin(int64(100), 200.0))

	assert.Equal(t, 100, findMin(100, 200))
	assert.Equal(t, 100.0, findMin(100, 200.0))
	assert.Equal(t, int64(100), findMin(int64(100), 200.0))
}

func getFirst[S []T, T any](slice S) T {
	data := slice[0]
	return data
}

func getSecond[S interface{ []T }, T any](slice S) T {
	data := slice[1]
	return data
}

func TestNestedTypeParameter(t *testing.T) {
	data := []string{"React", "Vue", "Svelte"}
	getFirst1 := getFirst[[]string, string](data)
	getFirst2 := getFirst[[]string](data)
	getSecond := getSecond(data)

	fmt.Println(getFirst1)
	fmt.Println(getFirst2)
	fmt.Println(getSecond)

	assert.Equal(t, "React", getFirst1)
	assert.Equal(t, "React", getFirst2)
	assert.Equal(t, "Vue", getSecond)
}
