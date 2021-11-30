package main

import "fmt"

func closure() func() {
	text := "Inside closure function"

	f := func() {
		fmt.Println(text)
	}

	return f
}

func main() {
	text := "Inside main"
	fmt.Println(text)

	newFunc := closure()
	newFunc()
}
