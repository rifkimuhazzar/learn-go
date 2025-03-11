package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found error")
	HelloError = errors.New("hello error")
)

func testError(value string) error {
	if value == "" {
		return ValidationError
	}  else if value == "b" {
		return HelloError
	} else if value != "a" {
		return NotFoundError
	} else {
		return nil
	}
}

func main() {
	result := testError("a")
	
	if result != nil {
		if errors.Is(result, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(result, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("uknown error")
		}
	} else {
		fmt.Println(result)
	}
}