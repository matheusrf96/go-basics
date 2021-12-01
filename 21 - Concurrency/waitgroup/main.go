package main

import (
	"fmt"
	"sync"
	"time"
)

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("Go routine")
		waitGroup.Done() // -1
	}()

	go func() {
		write("Hello eternity")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait()
	fmt.Println("Finish...")
}
