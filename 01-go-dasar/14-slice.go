package main

import "fmt"

func main() {
	months := [12]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 []string = months[4:6]
	fmt.Println(slice1)

	slice2 := months[:6]
	fmt.Println(slice2)

	slice3 := months[6:]
	fmt.Println(slice3)

	slice4 := months[:]
	fmt.Println(slice4)

	fmt.Println("---------------------------------------------------")

	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
		"Hello1",
	}
	
	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Saturday"
	daySlice1[1] = "Sunday"
	fmt.Println(daySlice1)
	fmt.Println(len(daySlice1))
	fmt.Println(cap(daySlice1))
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Holiday3")
	daySlice2[1] = "Holiday2"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(len(daySlice2))
	fmt.Println(cap(daySlice2))
	fmt.Println(days)
	
	fmt.Println("---------------------------------------------------")

	newSlice1 := make([]int, 2, 5)
	fmt.Println(newSlice1)
	fmt.Println(len(newSlice1))
	fmt.Println(cap(newSlice1))

	newSlice1[0] = 10
	newSlice1[1] = 20
	// newSlice[2] = 30 // error, harus menggunakan append()
	fmt.Println(newSlice1)

	newSlice2 := append(newSlice1, 30, 40)
	fmt.Println(newSlice1)
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = 100
	fmt.Println(newSlice1)
	fmt.Println(newSlice2)

	fmt.Println("---------------------------------------------------")

	fromSlice := 	days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	fmt.Println(fromSlice)
	fmt.Println(toSlice)
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	fmt.Println("---------------------------------------------------")

	thisIsArray := [...]int{1, 2, 3, 4, 5}
	thisIsSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(thisIsArray)
	fmt.Println(thisIsSlice)
}