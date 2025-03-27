package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes)
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("Hello World")
	logJson(100)
	logJson(true)
	logJson([]string{"React", "Vue", "Svelte"})
}
