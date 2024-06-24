package db

import (
	"crud/src/env"
	"database/sql"
	"log"
)

var CON *sql.DB

func Init() {
    cfg := env.LoadEnv()

    db, err := sql.Open(cfg.DBDriver, cfg.DBDatasource)
    if err != nil {
			log.Fatal(err)
    }

		CON = db
}

func SetDB(db *sql.DB) {
	CON = db
}