package generics

import (
	"fmt"
	"testing"
)

// 1
func sayHi[T any](value any) {
	var a T = value.(T)
	fmt.Println(a)
}

func Test1(t *testing.T) {
	sayHi[any]("Hi Svelte")
	sayHi[string]("Hi Svelte")
}

// 2
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

func printPersonDetails[T any](p T) {
	if person, ok := any(p).(Person); ok {
		fmt.Println("Name:", person.Name)
		fmt.Println("Age:", person.Age)
		person.Greet()
	} else {
		fmt.Println("Type is not Person")
	}
}

func Test2(t *testing.T) {
	p := Person{Name: "Svelte", Age: 25}
	printPersonDetails(p)
}

// 3
type Frontend struct {
	Name string
}

func (f Frontend) SayHello() string {
	return "Hello " + f.Name
}

type Backend struct {
	Name string
	// Age  uint8
}

func (b Backend) SayHello() string {
	return "Hello " + b.Name
}

type Job interface {
	Frontend | Backend
}

func Hello[T Frontend](value T) string {
	result := any(value).(Frontend).Name
	fmt.Println(result)
	return any(value).(Frontend).SayHello()
}

func Test3(t *testing.T) {
	frontend := Frontend{Name: "Svelte"}
	result1 := Hello(frontend)
	fmt.Println(result1)

	backend := Backend{Name: "Fiber"}
	result2 := Hello(Frontend(backend))
	fmt.Println(result2)
}
