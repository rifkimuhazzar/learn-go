package generics_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First T
	Last  T
}

func (d *Data[Z]) Say(name string) string {
	return "Hello " + name
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[Z]) ChangeFirst(first Z) Z {
	d.First = first
	return d.First
}

func TestGenericStruct(t *testing.T) {
	data := Data[string]{
		First: "Monkey D.",
		Last:  "Luffy",
	}
	fmt.Println(data)
	assert.Equal(t, "Hello Roronoa Zoro", data.Say("Roronoa Zoro"))
	assert.Equal(t, "Hello Roronoa Zoro", data.SayHello("Roronoa Zoro"))
	assert.Equal(t, "Roronoa Zoro", data.ChangeFirst("Roronoa Zoro"))
	fmt.Println(data)

	data2 := Data[any]{
		First: "Zoro",
		Last:  200,
	}
	fmt.Println(data2)
	assert.Equal(t, "Hello Roronoa Zoro", data2.Say("Roronoa Zoro"))
	assert.Equal(t, "Hello Roronoa Zoro", data2.SayHello("Roronoa Zoro"))
	assert.Equal(t, false, data2.ChangeFirst(false))
	fmt.Println(data2)
}
