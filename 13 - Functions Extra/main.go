package main

import "fmt"

var z int

// Init function runs before main function
func init() {
	fmt.Println("Init function running...")
	z = 5
}

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

// Recursive functions
func fib(pos uint) uint {
	if pos <= 1 {
		return pos
	}

	return fib(pos-2) + fib(pos-1)
}

// Defer
func func1() {
	fmt.Println("Exec func1")
}

func func2() {
	fmt.Println("Exec func2")
}

func studentIsApproved(n1 float32, n2 float32) (grade bool) {
	defer fmt.Println("Média calculada!")

	average := (n1 + n2) / 2

	if average >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println(z)
	// -----

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

	// -----

	fmt.Println("===== Fib =====")

	pos := uint(7)
	fmt.Println(fib(pos))

	for i := uint(1); i <= pos; i++ {
		fmt.Println(fib(i))
	}

	// -----

	fmt.Println("===== Defer =====")

	defer func1()
	func2()

	fmt.Println(studentIsApproved(9, 7))
}
