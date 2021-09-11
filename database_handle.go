package main

import (
	"database/sql"
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

	stmt, err := db.Prepare("INSERT INTO queries (query_text) VALUES (?)")
	if err != nil {
		// TODO: Should actually log to a file and return err
		log.Fatal("Unable to prepare statement:", err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(queryText)
	if err != nil {
		// TODO: Should actually log to a file and return err
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

	stmt, err := db.Prepare("INSERT INTO pages (domain, page_url, page_text, query_result_id) VALUES (?,?,?,?)")
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
