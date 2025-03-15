package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	AddToMap := func (value int) {
		defer group.Done()
		data.Store(value, value)
	}

	for i := range 10 {
		group.Add(1)
		go AddToMap(i)
	}

	group.Wait()
	data.Range(func(key, value any) bool {
		fmt.Print(key, ": ", value, "\n")
		return true
	})
	fmt.Println("TestMap function done!")
}