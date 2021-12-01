package main

import "fmt"

func reverseNumber(num int) int {
	return num * -1
}

func reverseNumberWithPointer(num *int) {
	*num = *num * -1
}

func main() {
	num := 23
	revNum := reverseNumber(num)
	fmt.Println(revNum)
	fmt.Println(num)

	num2 := 32
	fmt.Println(num2)
	reverseNumberWithPointer(&num2)
	fmt.Println(num2)
}
