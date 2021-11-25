package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number int8
}

func main() {
	address := address{street: "Nowhere", number: 5}

	var u user
	u.name = "Matheus"
	u.age = 25
	fmt.Println(u)

	u2 := user{"Matheus", 25, address}
	fmt.Println(u2)

	u3 := user{name: "Matheus"}
	fmt.Println(u3)
}
