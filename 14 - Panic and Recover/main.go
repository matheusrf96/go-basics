package main

import "fmt"

func recoverExec() {
	if r := recover(); r != nil {
		fmt.Println("Exec recovered!")
	}
}

func studentIsApproved(n1 float32, n2 float32) bool {
	defer recoverExec()

	avg := (n1 + n2) / 2

	if avg > 6 {
		return true
	} else if avg < 6 {
		return false
	}

	panic("THE AVERAGE IS SIX!!!")
}

func main() {
	fmt.Println(studentIsApproved(6, 6))
	fmt.Println("After function")
}
