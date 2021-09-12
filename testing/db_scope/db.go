package main

import (
	"database/sql"
	"log"
)

func QueryDB(queryUrl string) string {
	// CHECK IF ALREADY EXISTS, if so, throws error because I don't want to scrape that domain again
	var exists bool
	err := DB.QueryRow("SELECT EXISTS (SELECT * FROM query_results WHERE url=?)", queryUrl).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		// TODO: Return error here
		log.Fatal("test")
	}
	if exists {
		return "Exists"
	} else {
		return "Doesnt exists"
	}
}
