package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Drivers:", sql.Drivers())
	db, err := sql.Open("mysql", "root:7622446@tcp(127.0.0.1:3306)/web_scrape")
	if err != nil {
		log.Fatal("Unable to open connection to db")
	}
	defer db.Close()

	// SELECT
	results, err := db.Query("select * from pages")
	if err != nil {
		log.Fatal("Error when fetching product table rows:", err)
	}
	defer results.Close()
	for results.Next() {
		var (
			id   int
			page string
		)
		err = results.Scan(&id, &page)
		if err != nil {
			log.Fatal("Unable to parse row:", err)
		}
		fmt.Printf("ID: %d, Page: '%s'\n", id, page)
	}

	// INSERTS
	pages := []struct {
		page string
	}{
		{"teste 123"},
		{"teste 1234"},
		{"teste 12345"},
	}
	stmt, err := db.Prepare("INSERT INTO pages (page) VALUES (?)")
	if err != nil {
		log.Fatal("Unable to prepare statement:", err)
	}
	defer stmt.Close()
	for _, data := range pages {
		_, err = stmt.Exec(data.page)
		if err != nil {
			log.Fatal("Unable to execute statement:", err)
		}
	}

}
