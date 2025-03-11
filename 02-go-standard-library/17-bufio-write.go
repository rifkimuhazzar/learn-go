package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("Hello World 1\n")
	writer.WriteString("Hello World 2")
	writer.Flush()
}