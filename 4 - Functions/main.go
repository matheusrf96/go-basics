package main

import "fmt"

func sum(x int8, y int8) int8 {
	return x + y
}

func calcs(x, y int8) (int8, int8) {
	sum := x + y
	sub := x - y

	return sum, sub
}

func main() {
	sum := sum(5, 10)
	fmt.Println(sum)

	var f = func(txt string) string {
		return txt
	}
	fmt.Println(f("Function type"))

	f2 := func() {
		fmt.Println("Test")
	}
	f2()

	result_sum, result_sub := calcs(5, 8)
	fmt.Println(result_sum, result_sub)
}
