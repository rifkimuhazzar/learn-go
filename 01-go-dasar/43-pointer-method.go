package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	svelte := Man{"Svelte"}
	svelte.Married() // (&svelte).Married(), otomattis seperti ini
	fmt.Println(svelte)
	fmt.Println(svelte.Name)
}