package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving user data... (%s)\n", u.name)
}

func (u user) isAdult() bool {
	return u.age >= 18
}

func (u *user) happyBirthday() {
	u.age++
}

func main() {
	user1 := user{"User 1", 25}
	fmt.Println(user1)
	user1.save()

	user2 := user{"Matheus", 25}
	user2.save()

	isAdult := user2.isAdult()
	fmt.Println(isAdult)

	user2.happyBirthday()
	fmt.Println(user2.age)
}
