package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for range 100 {
		group.Add(1)
		go func() {
			defer group.Done()
			once.Do(increment)
		}()
	}
	
	group.Wait()
	fmt.Println("Counter:", counter)
}

func increment() {
	counter++
}

var counter int