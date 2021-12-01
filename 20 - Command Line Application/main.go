package main

import (
	"command-line-app/app"
	"fmt"
	"log"
	"os"
)

// Ex.: `go run main.go ip --host twitter.com`
func main() {
	fmt.Println("Starting point...")

	application := app.Generate()
	err := application.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
