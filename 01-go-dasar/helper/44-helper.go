package helper

import "fmt"

var version = "1.0.0"
var Aplication = "golang"

func sayGoodBye(name string) string {
	return "Good bye " + name
}

func Example(name string) string {
	fmt.Println(version)
	return sayGoodBye(name)
}

func SayHello(name string) string {
	return "Hello " + name
}
