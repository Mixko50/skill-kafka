package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Postgres(uri string) *sql.DB {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal("Fail to Connect to Database")
	}
	return db
}
