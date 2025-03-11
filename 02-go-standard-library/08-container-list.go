package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	fmt.Println(data)
	fmt.Println(*data)
	fmt.Println(&data)

	data.PushBack("React")
	data.PushBack("Vue")
	data.PushBack("Angular")
	data.PushBack("Svelte")
	data.PushBack("Svelte")
	data.PushBack("HTMX")
	
	fmt.Println("-----------------------")

	head := data.Front()
	fmt.Println(head)
	fmt.Println(head.Value)

	fmt.Println(head.Next())
	fmt.Println(head.Next().Value)
	fmt.Println(head.Next().Next().Next().Next().Next().Next())

	fmt.Println("-----------------------")

	for e := data.Front(); e != nil ; e = e.Next() {
		fmt.Println(e)
		fmt.Println(e.Value)
	}
}