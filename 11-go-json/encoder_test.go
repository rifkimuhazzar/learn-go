package gojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	file, _ := os.Create("./customer_encode.json")
	defer file.Close()

	encoder := json.NewEncoder(file)

	customer := &Customer{
		First: "React",
		Second: "Vue",
		Third: "Svelte",
	}
	encoder.Encode(customer)

	fmt.Println(customer)
}