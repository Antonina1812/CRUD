package db

import (
	"database/sql"
	"log"
)

var db sql.DB

func init() {
	var err error

	_, err = sql.Open("sqlite3", "movies.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS movies (
            id TEXT PRIMARY KEY,
            isbn TEXT,
            title TEXT,
            director_firstname TEXT,
            director_lastname TEXT
        )

	`)
	if err != nil {
		log.Fatal(err)
	}
}
