package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	googlesearch "github.com/rocketlaunchr/google-search"
)

var ctx = context.Background()

func main() {
	opts := googlesearch.SearchOptions{
		CountryCode:  "br",
		LanguageCode: "pt-br",
		Limit:        5,
	}

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
		fmt.Printf("Searching for: %s\n", queryText)

		returnLinks, err := googlesearch.Search(ctx, queryText, opts)
		if err != nil {
			fmt.Printf("something went wrong: %v", err)
			return
		}

		if len(returnLinks) == 0 {
			fmt.Printf("no results returned: %v", returnLinks)
		}

		fmt.Println("Results are: ", returnLinks)
		fmt.Println()
	}
}
