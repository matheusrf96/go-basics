package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	Host     = "localhost"
	Port     = "5432"
	User     = "postgres"
	Password = ""
	DbName   = "gobasics"
)

func Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		DbName,
		User,
		Password,
		Host,
		Port,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
