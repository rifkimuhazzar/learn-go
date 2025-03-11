package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Recover:", message)
}

func runApp(error bool)  {
	defer endApp()
	fmt.Println("----------")
	if error {
		panic("Panic")
	}
}

func main() {
	runApp(false)
	runApp(true)
	fmt.Println("--------------- func main() done")
}