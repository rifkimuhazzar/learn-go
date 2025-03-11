package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name 		string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email 	string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type Name: ", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		structField := valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"Vue"})

	fmt.Println("---------------------------------------------")

	readField(Person{"React", "USA", "react@example.com"})

	fmt.Println("---------------------------------------------")

	person := Person{
		Name: "a",
		Address: "a",
		Email: "a",
	}
	fmt.Println(IsValid(person))
}