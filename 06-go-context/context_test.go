package gocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextA.Value("b"))
}

func createCounter(ctx context.Context) <-chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <- ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	ctxParent := context.Background()
	ctx, cancel := context.WithCancel(ctxParent)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	destination := createCounter(ctx)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}

	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	ctxParent := context.Background()
	ctx, cancel := context.WithTimeout(ctxParent, 5 * time.Second)
	defer cancel()

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	destination := createCounter(ctx)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter:", n)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	ctxParent := context.Background()
	ctx, cancel := context.WithDeadline(ctxParent, time.Now().Add(5 * time.Second))
	defer cancel()

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	destination := createCounter(ctx)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter:", n)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}