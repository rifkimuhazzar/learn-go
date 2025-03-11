package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sumAll2(numbers [3]int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sumAll3(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sumAll(10))
	fmt.Println(sumAll(10, 10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10))
	
	numbers := []int{10, 10, 10}
	fmt.Println(sumAll(numbers...))
	
	numbers2 := [3]int{10, 10, 10}
	fmt.Println(sumAll(numbers2[:]...))
	fmt.Println(sumAll([]int{10, 10, 10}...))
	
	numbers3 := [3]int{10, 10, 10}
	fmt.Println(sumAll2(numbers3))
	fmt.Println(sumAll2([3]int{10, 10, 10}))
	fmt.Println(sumAll2([...]int{10, 10, 10}))
	
	fmt.Println(sumAll3([]int{10, 10, 10}))
}