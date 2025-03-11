package main

import "fmt"

type tech struct {
	frontEnd, backEnd string
}

func changFrontEndToSvelte(tech *tech) {
	tech.frontEnd = "Svelte"
}

func main() {
	// tech := tech{"React", "Express"}
	// fmt.Println(tech)
	// changFrontEndToSvelte(&tech)
	// fmt.Println(tech)

	tech := &tech{"React", "Express"}
	fmt.Println(tech)
	changFrontEndToSvelte(tech)
	fmt.Println(tech)
}