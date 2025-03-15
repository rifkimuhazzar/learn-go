package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T)  {
	group := &sync.WaitGroup{}
	for range 10 {
		group.Add(1)
		go RunAsynchronous(group)
	}
	group.Wait()
	fmt.Println("TestWaitGroup function done!")
}