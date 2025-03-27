package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {
	customer := Customer{
		First:   "React",
		Second:  "Vue",
		Third:   "Svelte",
		Hobbies: []string{"Coding", "Reading", "Watching Series/Anime"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"First":"React","Second":"Vue","Third":"Svelte","IsFrontEnd":false,"Number":0,"Hobbies":["Coding","Reading","Watching Series/Anime"]}`
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
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplexEncode(t *testing.T) {
	customer := Customer{
		First: "React",
		Addresses: []Address{
			{
				Street: "Street 1",
				Country: "Country 1",
				PostalCode: "PostalCode 1",
			},
			{
				Street: "Street 2",
				Country: "Country 2",
				PostalCode: "PostalCode 2",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"First":"React","Addresses":[{"Street":"Street 1","Country":"Country 1","PostalCode":"PostalCode 1"},{"Street":"Street 2","Country":"Country 2","PostalCode":"PostalCode 2"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexEncode(t *testing.T) {
	addresses := []Address{
		{
			Street: "Street 1",
			Country: "Country 1",
			PostalCode: "PostalCode 1",
		},
		{
			Street: "Street 2",
			Country: "Country 2",
			PostalCode: "PostalCode 2",
		},
	}

	bytes, err := json.Marshal(addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Street 1","Country":"Country 1","PostalCode":"PostalCode 1"},{"Street":"Street 2","Country":"Country 2","PostalCode":"PostalCode 2"}]`
	jsonBytes := []byte(jsonString)

	Addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, Addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(Addresses)
}
