package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONToStruct(t *testing.T) {
	jsonString := `{"First":"React","Second":"Vue","Third":"Svelte","IsFrontEnd":true,"Number":100}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.First)
	fmt.Println(customer.Second)
	fmt.Println(customer.Third)
}