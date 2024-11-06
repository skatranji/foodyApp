package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://user:password@localhost/foodyDb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
