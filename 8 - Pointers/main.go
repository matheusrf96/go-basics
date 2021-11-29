package main

import "fmt"

func main() {
	var num1 int = 5
	var num2 int = num1

	fmt.Println(num1, num2)

	num1++
	fmt.Println(num1, num2)

	// Pointers are memory references

	var num int
	var numPointer *int

	num = 10
	numPointer = &num

	fmt.Println(num, numPointer)
	fmt.Println(num, *numPointer)

	num = 13

	fmt.Println(num, numPointer)
	fmt.Println(num, *numPointer)
}
