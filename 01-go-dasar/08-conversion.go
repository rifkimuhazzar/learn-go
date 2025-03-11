package main

import "fmt"

func main() {
	var nilai32 int32 = 32767
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	name := "Hello World"
	e := name[1]
	eString := string(e)

	fmt.Println("name =", name)
	fmt.Println("e =", e)
	fmt.Println("eString =", eString)
	fmt.Println("string(101) =", string(101))
}