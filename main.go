package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open file of queries to Google
	file, err := os.Open("queries.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Loop over the lines in the file
	queries := bufio.NewScanner(file)

	for queries.Scan() {
		queryText := queries.Text()

		// Add query to database, get id
		queryId, err := AddQuery(queryText)
		if err != nil {
			// TODO: This should log fail query log
			fmt.Println("Failed to add query to database")
			continue
		}

		// Get top results of google
		queryResults, err := GoogleSearch(queryText)
		if err != nil {
			// TODO: Log to file and continue
			fmt.Println("Error getting google search results")
			continue
		}

		// Loop on results, add them to database, and send them to web crawler
		for _, res := range queryResults {
			// res is a struct with fields Description, Rank, Title, URL
			// Add query result to database, get query id
			resultId, err := AddQueryResults(res, queryId)
			if err != nil {
				// TODO: Log to file, if error, continue
				fmt.Println("Failed to add query result to database")
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
