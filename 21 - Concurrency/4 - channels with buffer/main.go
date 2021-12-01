package main

import "fmt"

func main() {
	c := make(chan string, 2)
	c <- "hi"
	c <- "hola"
	// As we setted the buffer as 2 if we set or get a third value to the channel, we'll have a deadlock

	msg := <-c
	msg2 := <-c
	fmt.Println(msg)
	fmt.Println(msg2)
}
