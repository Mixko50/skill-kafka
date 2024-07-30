package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Postgres(uri string) *sql.DB {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal("Fail to Connect to Database")
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Fail to Ping Database")
	}
	fmt.Println("Connected to Postgres db")
	return db
}
