package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// TODO: Change all print with file log.
func main() {
	// Open file of queries to Google
	file, err := os.Open("queries.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	queries := bufio.NewScanner(file)

	// Loop over the lines in the file
	for queries.Scan() {
		queryText := queries.Text()

		// Add query to database, get id
		queryId, err := AddQuery(queryText)
		if err != nil {
			fmt.Println("Add query error ", err)
			continue
		}

		// Get top results of google
		queryResults, err := GoogleSearch(queryText)
		// If no search result just go continue looping
		if err != nil {
			fmt.Println("Error getting google search results ", err)
			continue
		}

		// Loop on results, add them to database, and send them to web crawler
		for _, res := range queryResults {
			// res is a struct with fields {Description, Rank, Title, URL}
			// Add query result to database, get query id
			resultId, err := AddQueryResults(res, queryId)
			if err != nil {
				fmt.Println("Error adding query results ", err)
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
