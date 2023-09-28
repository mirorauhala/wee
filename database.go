package main

import (
	"database/sql"
	"fmt"
)

func setupDatabase() (*sql.DB, error) {
	dbURL := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the PostgreSQL database!")

	return db, nil
}

// Get the value of a key from the database
func getKey(dbConnection *sql.DB, key string) (string, error) {
	var value string
	err := dbConnection.QueryRow("SELECT value FROM kv WHERE key = $1 LIMIT 1", key).Scan(&value)
	if err != nil {
		return "", err
	}

	return value, nil
}
