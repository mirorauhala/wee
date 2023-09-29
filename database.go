package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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
