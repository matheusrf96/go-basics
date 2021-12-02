package main

import (
	"fmt"
	"math/rand"
	"time"
)

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}

func mux(c1 <-chan string, c2 <-chan string) <-chan string {
	output := make(chan string)

	go func() {
		for {
			select {
			case data := <-c1:
				output <- data
			case data := <-c2:
				output <- data
			}
		}
	}()

	return output
}

func main() {
	channel := mux(write("text a"), write("text b"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}
