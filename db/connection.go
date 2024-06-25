package db

import (
	"crud/env"
	"database/sql"
	"fmt"
	"log"

	"github.com/pressly/goose/v3"
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

func Migrate() error {
	fmt.Println("starting migrations")
	cfg := env.LoadEnv()

	driver := cfg.DBDriver
	dataSource := cfg.DBDatasource
	migrationsDir := cfg.MigrationPath

	db, err := sql.Open(driver, dataSource)
	if err != nil {
			return err
	}
	defer db.Close()

	fmt.Println("ending migrations")
	return goose.Up(db, migrationsDir)
}

func Rollback() error {
	cfg := env.LoadEnv()

	driver := cfg.DBDriver
	dataSource := cfg.DBDatasource
	migrationsDir := "migrations"

	db, err := sql.Open(driver, dataSource)
	if err != nil {
			return err
	}
	defer db.Close()

	return goose.Down(db, migrationsDir)
}