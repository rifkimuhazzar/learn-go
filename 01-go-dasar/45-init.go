package main

import (
	"01-go-dasar/database"
	_ "01-go-dasar/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}