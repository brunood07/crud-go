package db

import (
	"database/sql"
	"log"
)

var CON *sql.DB

func Init() {
	var err error
	CON, err = sql.Open("postgres", "postgres://admin:admin@localhost:5432/learning-go?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func SetDB(db *sql.DB) {
	CON = db
}