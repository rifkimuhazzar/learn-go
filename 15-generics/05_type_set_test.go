package generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	~int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		string
}

func min[T Number](value1, value2 T) T {
	if value1 < value2 {
		return value1
	} else {
		return value2
	}
}

type Age = int
type Age2 int

func TestMin(t *testing.T) {
	result := min(int8(100), int8(90))
	a, b := fmt.Println(result)
	fmt.Println(a)
	fmt.Println(b)
	println(a)
	println(b)

	assert.Equal(t, int8(90), min(int8(100), int8(90)))
	assert.Equal(t, 100, min(100, 190))
	assert.Equal(t, 100.0, min(100.0, 190.0))
	assert.Equal(t, float32(100.0), min[float32](100.0, 190.0))
	assert.Equal(t, "Svelte", min("Svelte", "Zve"))

	assert.Equal(t, 25, min[Age](25, 26))
	assert.Equal(t, Age(25), min[Age](25, 26))
	// assert.Equal(t, Age2(25), min[Age](25, 26))

	// assert.Equal(t, 25, min[Age2](25, 26))
	// assert.Equal(t, Age(25), min[Age2](25, 26))
	assert.Equal(t, Age2(25), min[Age2](25, 26))
}

func TestTypeInference(t *testing.T) {
	assert.Equal(t, 100, min(100, 200))
	assert.Equal(t, int16(100), min(int16(100), int16(200)))
	assert.Equal(t, 100.0, min(100.0, 200.0))
	assert.Equal(t, float32(100.0), min(float32(100), float32(200)))
	assert.Equal(t, float32(100), min[float32](100, 200))
	assert.Equal(t, 100, min(Age(100), Age(200)))
	assert.Equal(t, Age(100), min(Age(100), Age(200)))
	assert.Equal(t, Age2(100), min(Age2(100), Age2(200)))
}
