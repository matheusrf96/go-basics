package main

import "fmt"

func calcs(x int8, y int8) (sum int8, sub int8) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	sum, sub := calcs(5, 2)
	fmt.Println(sum, sub)
}
