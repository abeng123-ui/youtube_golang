package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Akbar"
		fmt.Println("Selesai baca channel")
	}()

	data := <-channel
	fmt.Print(data)
	time.Sleep(5 * time.Second)
}

func RunHelloWorld() {
	fmt.Println("Hello world")
}

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}
