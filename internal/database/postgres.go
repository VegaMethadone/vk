package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const connStr string = "user=postgres password=0000 dbname=testDB sslmode=disable"

func DBconnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
