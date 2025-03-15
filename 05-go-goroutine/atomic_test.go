package gogoroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	x := int32(0)
	group := sync.WaitGroup{}
	
	for i := 1; i <= 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			for j := 1; j <= 1000; j++ {
				atomic.AddInt32(&x, 1)
			}
		}()
	}

	group.Wait()
	fmt.Println(x)
}