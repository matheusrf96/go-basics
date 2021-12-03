package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type footballer struct {
	Name     string `json:"name"`
	Age      uint   `json:"age"`
	Position string `json:"position"`
	Club     string `json:"club"`
}

func main() {
	playerJSON := `{"name":"Renato Augusto","age":33,"position":"Midfielder","club":"Corinthians"}`

	var player footballer
	fmt.Println(player)

	err := json.Unmarshal([]byte(playerJSON), &player)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(player)
}
