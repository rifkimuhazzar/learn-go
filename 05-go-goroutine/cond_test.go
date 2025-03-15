package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = &sync.Mutex{}
var cond = sync.NewCond(locker)
var group = sync.WaitGroup{}

func WaitCondition(value int)  {
	defer group.Done()
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done:", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := range 10 {
		group.Add(1)
		go WaitCondition(i)
	}

	go func() {
		for range 10 {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// go func ()  {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()
	fmt.Println("TestCond function done!")
}