package gogoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	// defer close(channel)
	
	go func ()  {
		time.Sleep(time.Second * 3)
		channel <- "Hello World"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <- channel
	fmt.Println(data)
	time.Sleep(time.Second * 3)
	
	// close(channel)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	fmt.Println(<- channel)
	fmt.Println("TestChannelAsParameter function done!")
	time.Sleep(3 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(3 * time.Second)
	channel <- "Hello Go"
	fmt.Println("GiveMeResponse function done!")
}

func OnlyIn(channel chan<- string) {
	time.Sleep(3 * time.Second)
	channel <- "Hello Go"
	fmt.Println("OnlyIn function done!")
}

func OnlyOut(channel <-chan string) {
	time.Sleep(3 * time.Second)
	data := <- channel
	fmt.Println(data)
	fmt.Println("OnlyOut function done!")
}

func TestInAndOut(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)
	
	fmt.Println("TestInAndOut function done!")
	time.Sleep(4 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	fmt.Println(channel)
	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	channel <- "Hello Go"
	channel <- "Hello Go"
	channel <- "Hello Go"
	
	fmt.Println("TestBufferedChannel function done!")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func ()  {
		for i := range 10 {
			channel <- "Data - " + strconv.Itoa(i)
		}
		fmt.Println("goroutine 1 selesai!")
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Received:", data)
	}

	fmt.Println("goroutine main selesai!")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <- channel1:
			fmt.Println("channel1:", data)
			counter++
		case data := <- channel2:
			fmt.Println("channel2:", data)
			counter++
		}
		
		if counter == 2 {
			break
		}
	}

	fmt.Println("TestSelectChannel function done!")
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <- channel1:
			fmt.Println("channel1:", data)
			counter++
		case data := <- channel2:
			fmt.Println("channel2:", data)
			counter++
		default:
			fmt.Println("Menunggu data ...")
		}
		
		if counter == 2 {
			break
		}
	}

	fmt.Println("TestSelectChannel function done!")
}