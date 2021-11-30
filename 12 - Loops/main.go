package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		time.Sleep(time.Second)
		i++
		fmt.Println(i)
	}

	for j := 0; j < 10; j++ {
		time.Sleep(time.Second)
		fmt.Println(j)
	}

	codes := []string{"abc", "xyz", "mrf"}

	for key, value := range codes {
		fmt.Println(key, value)
	}

	for key, value := range "MATHEUS" {
		fmt.Println(key, string(value))
	}

	user := map[string]string{
		"name":    "Matheus",
		"surname": "Rodrigues",
	}

	for k, v := range user {
		fmt.Println(k, v)
	}

	for {
		fmt.Println("Endless loop")
		time.Sleep(time.Second)
	}
}
