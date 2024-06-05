package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbConnect() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:Dost0n1k@localhost:5432/book_store?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
