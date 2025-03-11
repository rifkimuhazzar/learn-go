package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`h([a-z])*o`)

	fmt.Println(regex.MatchString("ho"))
	fmt.Println(regex.MatchString("heo"))
	fmt.Println(regex.MatchString("Heo"))
	fmt.Println(regex.MatchString("hello"))
	fmt.Println(regex.MatchString("hello"))

	fmt.Println(regex.FindAllString("helo heo ho haa haha", 10))
}