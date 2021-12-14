package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Could not read request body"))
		return
	}

	var u user
	err = json.Unmarshal(body, &u)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Could not convert user into struct"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		w.Write([]byte("Could not insert data"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Data inserted succesfully. #%d", lastInsertId)))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Could not connect to database"))
		return
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Could not find users"))
		return
	}

	var users []user
	for rows.Next() {
		var user user

		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("Could not treat row found"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Could not encode users"))
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Could not convert param to uint"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Could not connect to database"))
		return
	}

	row, err := db.Query("SELECT * FROM users WHERE id = $1", ID)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Could not get data from database"))
		return
	}

	var user user
	if row.Next() {
		err := row.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("Could not scan user"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Could not encode user"))
		return
	}
}
