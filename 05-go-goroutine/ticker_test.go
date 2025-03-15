package gogoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		close(done)
	}()

	// for tick := range ticker.C {
	// 	fmt.Println(tick)
	// }

	loop:
		for {
			select {
			case tick := <- ticker.C:
				fmt.Println(tick)
			case <- done:
				fmt.Println("Keluar dari for!")
				break loop
			}
		}
	
	fmt.Println("TestTicker selesai!")
}

func TestTickChannel(t *testing.T) {
	channel := time.Tick(1 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		// channel.Stop()
		close(done)
	}()

	// for tick := range channel {
	// 	fmt.Println(tick)
	// }

	loop:
		for {
			select {
			case tick := <- channel:
				fmt.Println(tick)
			case <- done:
				fmt.Println("Keluar dari for!")
				break loop
			}
		}
	
	fmt.Println("TestTickChannel selesai!")
}