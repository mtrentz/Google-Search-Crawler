package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
)

// Global database variable
var DB *sql.DB

func main() {
	// Set up log file
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}
	log.SetOutput(logFile)

	// Connect to database, set it to global variable
	DB, err = sql.Open("mysql", "root:7622446@tcp(127.0.0.1:3306)/web_scrape")
	if err != nil {
		log.Fatal("Unable to open connection to db: ", err)
	}
	defer DB.Close()

	// Open file of queries to Google
	file, err := os.Open("queries.txt")
	if err != nil {
		log.Fatal("Error opening query file: ", err)
	}
	defer file.Close()

	queries := bufio.NewScanner(file)

	// Loop over the lines in the file
	for queries.Scan() {
		queryText := queries.Text()

		// Add query to database, get id
		queryId, err := AddQuery(queryText)
		if err != nil {
			log.Printf("Error adding query for %s. %s\n", queryText, err)
			continue
		}

		// Get top results of google
		queryResults, err := GoogleSearch(queryText)
		// If no search result just go continue looping
		if err != nil {
			log.Printf("Error getting google search results for %s. %s\n", queryText, err)
			continue
		}

		// Loop on results, add them to database, and send them to web crawler
		for _, res := range queryResults {
			// res is a struct with fields {Description, Rank, Title, URL}
			// Add query result to database, get query id
			resultId, err := AddQueryResults(res, queryId)
			if err != nil {
				log.Printf("Error adding query result for %s. %s\n", res.URL, err)
				// Continue on internal loop, over results
				continue
			}
			// Send to a recurisve crawl into that url domain
			CrawlURL(res.URL, resultId)
		}

	}
	if err := queries.Err(); err != nil {
		fmt.Println("Error scanning queries text file")
		log.Fatal(err)
	}
}
