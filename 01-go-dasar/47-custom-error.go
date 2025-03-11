package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	} else if id != "go" {
		return &notFoundError{"not found error"}
	} else {
		return nil
	}
}

func main() {
	err := SaveData("", nil)

	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("validation error:", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("not found error:", notFoundErr.Error())
		} else {
			fmt.Println("unknown error:", err.Error())
		}

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("unknown error:", finalError.Error())
		}
	} else {
		fmt.Println("Success")
	}

	var a any = 100
	fmt.Println(a)

	b, ok := a.(int)
	fmt.Println(b)
	fmt.Println(ok)

	c := int32(a.(int))
	fmt.Println(c)
	
	d, ok := a.(*int)
	fmt.Println(d)
	fmt.Println(ok)
}