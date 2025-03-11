package main

import "fmt"

func main() {
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Counter", counter)
	// 	counter++
	// }
	// fmt.Println("Selesai")

	// for counter := 1; counter <= 10; {
	// 	fmt.Println("Counter", counter)
	// 	counter++
	// }
	// fmt.Println("Selesai")
	
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Counter", counter)
	}
	fmt.Println("Selesai")

	names := [...]string{"react", "vue", "svelte"}
	for i := 0; i < len(names); i++ {
		fmt.Println("name =", names[i])
	}
	for i, name := range names {
		fmt.Println("i -", i, "=", name)
	}
	for _, name := range names {
		fmt.Println(name)
	}
}