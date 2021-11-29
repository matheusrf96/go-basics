package main

import "fmt"

func main() {
	// map[<key type>]<value type>
	user := map[string]string{
		"name":    "Matheus",
		"surname": "Rodrigues",
	}
	fmt.Println(user)
	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"first name": "name",
			"last name":  "test",
		},
		"course": {
			"name":   "History",
			"campus": "Campus II",
		},
	}
	fmt.Println(user2)
	fmt.Println(user2["course"]["name"])

	delete(user2, "name")
	fmt.Println(user2)

	user2["sign"] = map[string]string{
		"name": "Scorpio",
	}
	fmt.Println(user2)
}
