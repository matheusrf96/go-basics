package main

import "fmt"

func main() {
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 != 2)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	five := 5
	five++
	fmt.Println(five)
}
