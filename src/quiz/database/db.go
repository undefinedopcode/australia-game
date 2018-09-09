package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB    *sql.DB
	dberr error
)

func init() {
	DB, dberr = sql.Open("sqlite3", "../../data/australia.db")
	if dberr != nil {
		log.Fatal(dberr)
	}
}
