package main

import (
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("root"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("loading users page..."))
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/home", home)
	http.HandleFunc("/users", users)

	log.Println("Server running at: 0.0.0.0:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
