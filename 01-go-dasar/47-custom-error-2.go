package main

import (
	"fmt"
)

type ValidationError struct {
	message string
}

func (v *ValidationError) Error() string {
	return v.message
}

type NotFoundError struct {
	message string
}

func (n *NotFoundError) Error() string {
	return n.message
}

func testError(value string) error {
	if value == "" {
		return &ValidationError{"Terjadi validation error"}
	} else if value != "a" {
		return &NotFoundError{"Terjadi not found error"}
	} else {
		return nil
	}
}

func main() {
	error := testError("")
	fmt.Println(error)
	T, ok := error.(*ValidationError);
	fmt.Println(T)
	fmt.Println(ok)

	if error != nil {
		if T, ok := error.(*ValidationError); ok {
			fmt.Println("ValidationError:", T)
		} else if T, ok := error.(*NotFoundError); ok {
			fmt.Println("NotFoundError:",T)
		} else {
			fmt.Println("UnknownError:", error)
		}

		switch error := error.(type) {
		case *ValidationError:
			fmt.Println("ValidationError:", error)
		case *NotFoundError:
			fmt.Println("NotFoundError:", error)
		default:
			fmt.Println("UnknownError:", error)
		}
	} else {
		fmt.Println("Success")
	}
}