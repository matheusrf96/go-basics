package main

import (
	"fmt"
	"packages/auxpack"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hi...")
	auxpack.Write()

	err := checkmail.ValidateFormat("matheusrf96@gmail.com")
	fmt.Println(err)
}
