package gogoroutine

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Run: Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go HelloWorld()
	fmt.Println("Run: TestCreateGoroutine")
	time.Sleep(time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := range 100000 {
		go DisplayNumber(i)
	}
	time.Sleep(time.Second * 5)
}
