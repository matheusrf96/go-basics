package main

import (
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("root"))
}

func main() {
	http.HandleFunc("/", root)

	log.Println("Server running at: 0.0.0.0:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
