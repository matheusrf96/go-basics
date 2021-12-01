package main

import (
	"fmt"
	"time"
)

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	go write("Go routine")
	write("Hello eternity")
}
