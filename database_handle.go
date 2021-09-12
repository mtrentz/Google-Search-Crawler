package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	googlesearch "github.com/rocketlaunchr/google-search"
)

func AddQuery(queryText string) (id int64, err error) {
	// Open connection to database
	db, err := sql.Open("mysql", "root:7622446@tcp(127.0.0.1:3306)/web_scrape")
	if err != nil {
		log.Fatal("Unable to open connection to db")
	}
	defer db.Close()

	// CHECK IF ALREADY EXISTS
	var existingId int64
	err = db.QueryRow("SELECT id FROM queries WHERE query_text=?", queryText).Scan(&existingId)
	// If no rows exists would throw an error. In that case I don't want it to log error
	if err != nil && err != sql.ErrNoRows {
		// TODO: Return error and log
		log.Fatal("Error on select statement")
	}
	// If different from 0, query was already added, so return existing id
	// Else, continue normally and insert into database
	if existingId != 0 {
		return existingId, nil
	}

	// INSERT INTO DB
	insert_stmt, err := db.Prepare("INSERT INTO queries (query_text) VALUES (?)")
	if err != nil {
		// TODO: Return error and no id
		log.Fatal("Unable to prepare insert statement:", err)
	}
	defer insert_stmt.Close()
	res, err := insert_stmt.Exec(queryText)
	if err != nil {
		// TODO: Return error and no id
		log.Fatal("Unable to execute statement:", err)
	}
	lid, err := res.LastInsertId()

	return lid, nil
}

func AddQueryResults(queryResult googlesearch.Result, queryId int64) (id int64, err error) {
	// Open connection to database
	db, err := sql.Open("mysql", "root:7622446@tcp(127.0.0.1:3306)/web_scrape")
	if err != nil {
		//  TODO: Return error and log to file
		log.Fatal("Unable to open connection to db")
	}
	defer db.Close()

	// CHECK IF ALREADY EXISTS, if so, throws error because I don't want to scrape that domain again
	var exists bool
	testUrl := "https://www.maxiquim.com.br/"
	err = db.QueryRow("SELECT EXISTS (SELECT * FROM query_results WHERE url=?)", testUrl).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		// TODO: Return error here
		log.Fatal("test")
	}
	if exists {
		// TODO: Return error here
		fmt.Println("Exists")
	}

	stmt, err := db.Prepare("INSERT INTO query_results (result_rank, title, url, description, query_id) VALUES (?,?,?,?,?)")
	if err != nil {
		// TODO: Should actually log to a file and return err
		log.Fatal("Unable to prepare statement:", err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(queryResult.Rank, queryResult.Title, queryResult.URL, queryResult.Description, queryId)
	if err != nil {
		// TODO: Should actually log to a file and return err
		log.Fatal("Unable to execute statement:", err)
	}
	lid, err := res.LastInsertId()

	return lid, nil

}

func AddPage(pageText string, domain string, pageUrl string, resultId int64) {
	// Open connection to database
	db, err := sql.Open("mysql", "root:7622446@tcp(127.0.0.1:3306)/web_scrape")
	if err != nil {
		//  TODO: Return error and log to file
		log.Fatal("Unable to open connection to db")
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT IGNORE INTO pages (domain, page_url, page_text, query_result_id) VALUES (?,?,?,?)")
	if err != nil {
		// TODO: Should actually log to a file and return err
		log.Fatal("Unable to prepare statement:", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(domain, pageUrl, pageText, resultId)
	if err != nil {
		// TODO: Should actually log to a file and return err
		log.Fatal("Unable to execute statement:", err)
	}

}
