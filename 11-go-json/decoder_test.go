package gojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	file, _ := os.Open("./customer.json")
	defer file.Close()

	decoder := json.NewDecoder(file)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}