package gogoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func Test(t *testing.T) {
	group := sync.WaitGroup{}
	for range 100 {
		group.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			group.Done()
		}()
	}

	totalCPUs := runtime.NumCPU()
	fmt.Println("Total CPUs:", totalCPUs)
	
	totalThreads := runtime.GOMAXPROCS(0)
	fmt.Println("Total Threads:", totalThreads)
	
	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines:", totalGoroutines)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for range 100 {
		group.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			group.Done()
		}()
	}

	totalCPUs := runtime.NumCPU()
	fmt.Println("Total CPUs:", totalCPUs)
	
	totalThreads1 := runtime.GOMAXPROCS(25)
	fmt.Println("Total Threads1:", totalThreads1)
	totalThreads2 := runtime.GOMAXPROCS(0)
	fmt.Println("Total Threads2:", totalThreads2)
	totalThreads3 := runtime.GOMAXPROCS(0)
	fmt.Println("Total Threads2:", totalThreads3)
	
	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines:", totalGoroutines)

	group.Wait()
}