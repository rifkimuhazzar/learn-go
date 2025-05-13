package generics

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func experimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	result1 := experimentalMin(100, 200)
	result2 := experimentalMin("Svelte", "Zve")
	fmt.Println(result1)
	fmt.Println(result2)
	assert.Equal(t, 100, result1)
	assert.Equal(t, "Svelte", result2)
}

func TestMaps(t *testing.T) {
	data1 := map[string]string{"Name": "Svelte"}
	data2 := map[string]string{"Name": "Svelte"}
	assert.True(t, maps.Equal(data1, data2))
}

func TestSlices(t *testing.T) {
	data1 := []string{"Svelte"}
	data2 := []string{"Svelte"}
	assert.True(t, slices.Equal(data1, data2))
}
