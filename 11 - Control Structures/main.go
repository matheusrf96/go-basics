package main

import "fmt"

func weekday(num int) string {
	// Switch case works as `case var == value:`
	// ¹`fallthrough`` jumps to next case

	switch num {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid day :("
	}
}

func conditionals(num int) {
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

func main() {
	num := 10
	conditionals(num)

	day := weekday(5)
	fmt.Println(day)
}
