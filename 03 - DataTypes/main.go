package main

import (
	"errors"
	"fmt"
)

func main() {
	var num int16 = -10000
	fmt.Println(num)

	var num2 uint32 = 5
	fmt.Println(num2)

	// alias
	// int32 == rune
	var num3 rune = 12345
	fmt.Println(num3)

	// uint8 == byte
	var num4 byte = 123
	fmt.Println(num4)

	var num5 float32 = 123.45
	fmt.Println(num5)

	var num6 float64 = 321.1232234123412412421424
	fmt.Println(num6)

	num7 := 12345.67
	fmt.Println(num7)

	var str string = "abc"
	fmt.Println(str)

	char := 'M'
	fmt.Println(char)

	var text string
	fmt.Println(text)

	var null int8
	fmt.Println(null)

	var boolean1 bool
	fmt.Println(boolean1)

	var err1 error
	fmt.Println(err1)

	var err2 error = errors.New("Internal Error :(")
	fmt.Println(err2)
}
