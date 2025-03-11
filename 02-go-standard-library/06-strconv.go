package main

import (
	"fmt"
	"strconv"
)

func main() {
	res, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(res)
		fmt.Println(err)
	} else {
		fmt.Println(res)
		fmt.Println(err)
	}
	
	fmt.Println("------------------------------------------")

	// resInt, errInt := strconv.ParseInt("100", 10, 16)
	resInt, errInt := strconv.Atoi("100")
	if err != nil {
		fmt.Println(resInt)
		fmt.Println(errInt)
	} else {
		fmt.Println(resInt)
		fmt.Println(errInt)
	}

	binary := strconv.FormatInt(190, 2)
	fmt.Println(binary)
	
	stringInt := strconv.Itoa(180) 
	fmt.Println(stringInt)
}