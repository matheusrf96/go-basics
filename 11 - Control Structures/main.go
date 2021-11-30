package main

import "fmt"

func main() {
	num := 10

	if num > 5 {
		fmt.Println("Greater than 5")
	} else {
		fmt.Println("Less or equal 5")
	}

	if num2 := num; num2 > 0 {
		fmt.Println("Num2 is greater than zero.")
	} else if num2 < -10 {
		fmt.Println("Less than -10")
	} else {
		fmt.Println("Between -10 and 0")
	}
}
