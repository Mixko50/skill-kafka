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
	fmt.Println("Connected to Postgres")
	return db
}
