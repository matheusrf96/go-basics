package main

import "fmt"

// Named return
func calcs(x int8, y int8) (sum int8, sub int8) {
	sum = x + y
	sub = x - y
	return
}

// Variadic functions
func openSum(nums ...int) int {
	result := 0

	for _, n := range nums {
		result += n
	}

	return result
}

func writeNums(text string, nums ...int) {
	for _, num := range nums {
		fmt.Println(text, num)
	}
}

func main() {
	sum, sub := calcs(5, 2)
	fmt.Println(sum, sub)

	// -----

	sum2 := openSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(sum2)

	writeNums("Olá", 3, 2, 1)

	// Anonymous function

	anon := func(text string) string {
		return fmt.Sprintf("Received -> %s", text)
	}("Text example")

	fmt.Println(anon)
}
