package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	writer.Write([]string{"React 1", "Vue 1", "Svelte 1"})
	writer.Write([]string{"React 2", "Vue 2", "Svelte 2"})
	writer.Write([]string{"React 3", "Vue 3", "Svelte 3"})

	writer.Flush()
}