package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street 		 string
	Country		 string
	PostalCode string
}

type Customer struct {
	First, Second, Third string
	IsFrontEnd bool
	Number int
	Hobbies []string
	Addresses []Address
}

func TestStructToJSON(t *testing.T) {
	customer := Customer{
		First: "React", 
		Second: "Vue", 
		Third: "Svelte", 
		IsFrontEnd: true, 
		Number: 100,
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes)
	fmt.Println(string(bytes))
}