package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err  := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
		fmt.Println(err)
	} else {
		fmt.Println(hostname)
		fmt.Println(err)
	}
}