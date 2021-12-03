package main

import (
	"bytes"
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
	player := footballer{"Renato Augusto", 33, "Midfielder", "Corinthians"}
	fmt.Println(player)

	playerJSON, err := json.Marshal(player)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(playerJSON)
	fmt.Println(bytes.NewBuffer(playerJSON))
}
