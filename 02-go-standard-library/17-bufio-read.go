package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("This is first line\nThis is second line\nThis is third line")
	reader := bufio.NewReader(input)

	for {
		var line, prefix, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(line)
		fmt.Println(string(line))
		fmt.Println(prefix)
		fmt.Println(err)
		fmt.Println("--------------")
	}
}