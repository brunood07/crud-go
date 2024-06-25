package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DBDriver   string
	DBDatasource string
	Port       string
	MigrationPath string
}

func LoadEnv() *Env {
	err := godotenv.Load()
	if err != nil {
			log.Fatal("Error loading .env file")
	}

	DBDriver := os.Getenv("DRIVER_NAME")
	DBDatasource := os.Getenv("DATASOURCE")
	Port := os.Getenv("PORT")
	MigrationPath := os.Getenv("MIGRATION_PATH")

	return &Env{
		DBDriver:   DBDriver,
		DBDatasource: DBDatasource,
		Port:       Port,
		MigrationPath: MigrationPath,
	}
}