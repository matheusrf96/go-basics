package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Could not read request body"))
		return
	}

	var u user
	err = json.Unmarshal(body, &u)
	if err != nil {
		w.Write([]byte("Could not convert user into struct"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Could not connect to database"))
		return
	}

	statement, err := db.Prepare(`INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Could not create statement"))
		return
	}
	defer statement.Close()

	lastInsertId := 0
	err = statement.QueryRow(u.Name, u.Email).Scan(&lastInsertId)
	if err != nil {
		w.Write([]byte("Could not insert data"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Data inserted succesfully. #%d", lastInsertId)))
}
