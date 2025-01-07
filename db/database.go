package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SSLMODE"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Create the database if it does not exist
	if err := createDatabase(db); err != nil {
		return nil, err
	}

	// Close the initial connection before connecting to the specific database
	db.Close()

	// Connect to the specific database
	connStr = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))
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
