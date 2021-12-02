package main

import (
	"fmt"
	"time"
)

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}

func main() {
	channel := write("hi")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}
