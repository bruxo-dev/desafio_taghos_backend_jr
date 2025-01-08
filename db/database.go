package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SSLMODE"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := createDatabase(db); err != nil {
		return nil, err
	}

	db.Close()

	connStr = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

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
