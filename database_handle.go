package main

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	googlesearch "github.com/rocketlaunchr/google-search"
)

var ErrSqlSelect = errors.New("ErrSqlSelect: error querying the database")
var ErrSqlPrepare = errors.New("ErrSqlPrepare: error preparing sql statement")
var ErrSqlInsert = errors.New("ErrSqlInsert: error inserting into database")
var ErrRowExists = errors.New("ErrRowExists: error inserting duplicate value")

func AddQuery(queryText string) (id int64, err error) {
	// CHECK IF ALREADY EXISTS
	var existingId int64
	err = DB.QueryRow("SELECT id FROM queries WHERE query_text=?", queryText).Scan(&existingId)
	// If no rows exists it will throw an ErrNoRows. In that case I won't consider as actual error
	if err != nil && err != sql.ErrNoRows {
		return 0, ErrSqlSelect
	}
	// If different from 0, query was already added, so return existing id
	if existingId != 0 {
		return existingId, nil
	}

	// INSERT INTO DB
	insert_stmt, err := DB.Prepare("INSERT INTO queries (query_text) VALUES (?)")
	if err != nil {
		return 0, ErrSqlPrepare
	}
	defer insert_stmt.Close()

	res, err := insert_stmt.Exec(queryText)
	if err != nil {
		return 0, ErrSqlInsert
	}
	lid, err := res.LastInsertId()

	return lid, nil
}

func AddQueryResults(queryResult googlesearch.Result, queryId int64) (id int64, err error) {
	// CHECK IF ALREADY EXISTS, if so, throws error because I don't want to scrape that domain again
	var exists bool
	err = DB.QueryRow("SELECT EXISTS (SELECT * FROM query_results WHERE url=?)", queryResult.URL).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return 0, ErrSqlSelect
	}
	if exists {
		// On main i'll just check for any error and continue to next domain scrape
		return 0, ErrRowExists
	}

	stmt, err := DB.Prepare("INSERT INTO query_results (result_rank, title, url, description, query_id) VALUES (?,?,?,?,?)")
	if err != nil {
		return 0, ErrSqlPrepare
	}
	defer stmt.Close()
	res, err := stmt.Exec(queryResult.Rank, queryResult.Title, queryResult.URL, queryResult.Description, queryId)
	if err != nil {
		return 0, ErrSqlInsert
	}
	lid, err := res.LastInsertId()

	return lid, nil

}

func AddPage(pageText string, domain string, pageUrl string, resultId int64) error {
	stmt, err := DB.Prepare("INSERT IGNORE INTO pages (domain, page_url, page_text, query_result_id) VALUES (?,?,?,?)")
	if err != nil {
		return ErrSqlPrepare
	}
	defer stmt.Close()
	_, err = stmt.Exec(domain, pageUrl, pageText, resultId)
	if err != nil {
		return ErrSqlInsert
	}

	return nil
}
