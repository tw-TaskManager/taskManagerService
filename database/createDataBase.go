package database

import "database/sql"

const (
	DB_DRIVER = "postgres"
	DB_CONNECTION = "user=postgres dbname=postgres password=postgres sslmode=disable"
)

func OreateDataBase() (*sql.DB, error) {
	db, err := sql.Open(DB_DRIVER, DB_CONNECTION)
	return db, err;
}