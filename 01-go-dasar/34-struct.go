package main

import "fmt"

// type Customer struct {
// 	Name    string
// 	Address string
// 	Age     int
// }

type Customer struct {
	Name, Address string
	Age 					int
	active 				bool
}

func (x Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", x.Name)
}

func main() {
	var person Customer
	person.Name = "React"
	person.Address = "USA"
	person.Age = 10
	fmt.Println(person)
	fmt.Println(person.Name)
	fmt.Println(person.Address)
	fmt.Println(person.Age)
	
	fmt.Println("------------------")
	
	person2 := Customer{
		Name: "Vue",
		Age: 10,
		Address: "USA",
	}
	fmt.Println(person2)
	
	person3 := Customer{"Svelte", "USA", 10, false}
	fmt.Println(person3)

	person.sayHello("Test")
	person2.sayHello("Test")
	person3.sayHello("Test")
}