package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := "user=postgres password=123123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Create the database if it does not exist
	if err := createDatabase(db); err != nil {
		return nil, err
	}

	// Connect to the specific database
	connStr = "user=postgres password=123123 dbname=desafio_taghos sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Initialize the database schema
	if err := initDB(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createDatabase(db *sql.DB) error {
	_, err := db.Exec("CREATE DATABASE desafio_taghos")
	if err != nil && err.Error() != `pq: database "desafio_taghos" already exists` {
		return err
	}
	return nil
}

func initDB(db *sql.DB) error {
	file, err := os.ReadFile("db/init.sql")
	if err != nil {
		return err
	}
	if _, err := db.Exec(string(file)); err != nil {
		return err
	}
	return nil
}
