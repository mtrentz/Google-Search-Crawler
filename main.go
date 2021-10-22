package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

// Global database variable
var DB *sql.DB
var wg sync.WaitGroup

func connectDB() {

	var err error

	fmt.Println("Connecting to database...")
	func() {
		retries := 10
		count := 0

		for {
			// Connect to database, set it to global variable
			// TODO: Get OS Env variables from Docker Compose
			DB, err = sql.Open("mysql", "crawler:crawler@tcp(crawler_db:3308)/crawler")
			if err != nil {
				fmt.Println("Unable to open connection to db: ", err)
			} else {
				// If connected, exits function
				return
			}
			fmt.Println("Trying again in 5 seconds.")
			time.Sleep(time.Second * 5)
			count++

			if count >= retries {
				log.Fatal("Could not connect to database, exiting: ", err)
				return
			}
		}
	}()
}

func main() {
	// Set up log file
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}
	log.SetOutput(logFile)

	connectDB()

	// Open file of queries to Google
	fmt.Println("Reading queries...")
	file, err := os.Open("queries.txt")
	if err != nil {
		log.Fatal("Error opening query file: ", err)
	}
	defer file.Close()

	queries := bufio.NewScanner(file)
	// Get num of lines in file (passes file name to read file again)
	queriesCount := countFileLines("queries.txt")
	fmt.Printf("Started scraping all %d google queries\n", queriesCount)

	// Preparing setup for concurrency
	maxGoroutines := 5
	guard := make(chan int, maxGoroutines)
	wg.Add(queriesCount)

	// Load unallowed domains from file
	var unallowedDomainsList = loadUnallowedDomainList()

	// Loop over the lines in the file
	for queries.Scan() {
		queryText := queries.Text()

		// would block if guard channel is already filled
		// guard <- struct{}{}
		guard <- 1 // will block if there is maxGoroutines ints in sem

		// Spin goroutines for each google query
		go func() {
			// Add query to database, get id
			queryId, err := AddQuery(queryText)
			if err != nil {
				log.Printf("Error adding query for %s. %s\n", queryText, err)
				return
			}

			// Get top results of google
			queryResults, err := GoogleSearch(queryText)
			// If no search result just go continue looping
			if err != nil {
				log.Printf("Error getting google search results for %s. %s\n", queryText, err)
				return
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

				// Check if domain not on unallowed list
				domain, err := getDomain(res.URL)
				if err != nil {
					return
				}
				if stringNotInSlice(domain, unallowedDomainsList) {
					// Send to a recurisve crawl into that url domain
					CrawlURL(res.URL, resultId)
				}

			}
			wg.Done()
			// Sends back info knowing it stopped a goroutine
			<-guard
		}()

	}
	if err := queries.Err(); err != nil {
		fmt.Println("Error scanning queries text file")
		log.Fatal(err)
	}

	wg.Wait()
}

func countFileLines(fileName string) (n int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening query file: ", err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	var lineCount int
	for s.Scan() {
		lineCount++
	}
	return lineCount
}

func loadUnallowedDomainList() []string {
	file, err := os.Open("unallowed_domains.txt")
	if err != nil {
		log.Fatal("Error opening unallowed domains list: ", err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	var unallowedDomains []string
	for s.Scan() {
		unallowedDomains = append(unallowedDomains, s.Text())
	}
	return unallowedDomains
}

// Function to check if string is in NOT slice of strings
func stringNotInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return false
		}
	}
	return true
}

// Function to get the domain of a URL
func getDomain(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		log.Printf("Error parsing domain of url: %s", err)
		return "", err
	}
	parts := strings.Split(u.Hostname(), ".")
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]

	return domain, nil
}
