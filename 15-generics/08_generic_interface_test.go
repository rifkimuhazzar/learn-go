package generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	SetValue(value T)
	GetValue() T
}

func changeValue[T any](parameter GetterSetter[T], value T) T {
	parameter.SetValue(value)
	return parameter.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (m *MyData[T]) SetValue(value T) {
	m.Value = value
}

func (m *MyData[T]) GetValue() T {
	return m.Value
}

func TestGenericInterface(t *testing.T) {
	myData := MyData[string]{}
	result := changeValue(&myData, "Monkey D. Luffy")
	fmt.Println(result)

	assert.Equal(t, "Monkey D. Luffy", myData.Value)
	assert.Equal(t, "Monkey D. Luffy", result)
}
