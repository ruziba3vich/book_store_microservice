package main

import (
	"authors/api"
	"authors/storage"
	"authors/storage/db"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := db.DbConnect()
	author := storage.NewAuthor(db)
	r := api.Router(author)
	fmt.Println("server is listening on 7777")
	err := http.ListenAndServe(":7777", r)
	if err != nil {
		log.Fatal(err)
	}
}
