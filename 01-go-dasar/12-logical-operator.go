package main

import "fmt"

func main() {
	nilaiAkhir := 90
	nilaiKehadiran := 81
	isPass := nilaiAkhir > 80 && nilaiKehadiran > 80

	fmt.Println(isPass)
}