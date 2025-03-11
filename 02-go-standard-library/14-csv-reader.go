package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "React 1, Vue 1, Svelte 1\n" + "React 2, Vue 2, Svelte 2\n" + "React 3, Vue 3, Svelte 3"
	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}