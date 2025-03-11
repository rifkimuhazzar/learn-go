package main

import "fmt"

func main() {
	// person := map[string]string{}
	// person["name1"] = "React"
	// person["name2"] = "Vue"
	// fmt.Println(person)

	person := map[string]string{
		"name1": "React",
		"name2": "Vue",
	}
	fmt.Println(person)
	fmt.Println(person["name1"])
	fmt.Println(person["name2"])

	name := make(map[string]string)
	name["library"] = "React"
	name["framework"] = "Next"
	name["other"] = "Express"
	fmt.Println(name)
	delete(name, "other")
	fmt.Println(name)
}