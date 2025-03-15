package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			s := "Pool is empty"
			return &s
		},
	}

	a := "react"
	b := "vue"
	c := "100"

	pool.Put(&a)
	pool.Put(&b)
	pool.Put(&c)

	for range 5 {
		go func() {
			data := pool.Get().(*string)
			fmt.Println(*data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("TestPool function done!")
}
