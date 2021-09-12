package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gocolly/colly"
)

func CrawlURL(resultUrl string, resultId int64) {
	// Set a max amount of web pages that it should go through when crawling the resultUrl
	// When exceeded, will panic and recover here, exiting the web crawler for this url.
	const MAXSCRAPE = 5
	const MAXVISIT = 20
	var scrapeCounter int
	var visitCounter int
	// Recover
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Crawler exited when crawling base url %s\n", resultUrl)
			return
		}
	}()

	// Parse URL
	u, err := url.Parse(resultUrl)
	if err != nil {
		// If cant parse URL, just exists function
		log.Printf("Error parsing url for: %s\n", resultUrl)
		return
	}
	domain := u.Hostname()

	c := colly.NewCollector()
	c.MaxDepth = 5
	c.AllowedDomains = []string{domain}
	c.AllowURLRevisit = false
	// c.Async = true

	c.OnHTML("html body", func(e *colly.HTMLElement) {
		// Page URL (since it recursively goes into all hrefs)
		pageUrl := e.Request.URL.String()
		// Get the HTML string
		pageHtml, _ := e.DOM.Html()
		// Parse only the text
		pageText, err := ParseHTML(pageHtml)
		if err != nil {
			// This will exit the crawler
			panic("Exit")
		}
		err = AddPage(pageText, domain, pageUrl, resultId)
		if err != nil {
			fmt.Printf("Error adding page into database for %s. %s\n", pageUrl, err)
		}

		// Checks if already maxed out scraped per domain, if so, exits crawler
		scrapeCounter++
		// Check here also if already maxed out visits or scrapes, if so, exits crawler
		if scrapeCounter >= MAXSCRAPE || visitCounter >= MAXVISIT {
			panic("Exit")
		}
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Check here also if already maxed out visits or scrapes, if so, exits crawler
		if scrapeCounter >= MAXSCRAPE || visitCounter >= MAXVISIT {
			panic("Exit")
		}
		visitCounter++
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		// Check here also if already maxed out visits or scrapes, if so, exits crawler
		if scrapeCounter >= MAXSCRAPE || visitCounter >= MAXVISIT {
			panic("Exit")
		}
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(resultUrl)

	// c.Wait()
}
