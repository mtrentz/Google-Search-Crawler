package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gocolly/colly"
)

func CrawlURL(resultUrl string, resultId int64) {
	// Parse URL
	u, err := url.Parse(resultUrl)
	if err != nil {
		log.Fatal(err)
	}
	domain := u.Hostname()

	c := colly.NewCollector()
	c.MaxDepth = 2
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
			log.Panic(err)
		}
		AddPage(pageText, domain, pageUrl, resultId)

	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(resultUrl)

	// c.Wait()
}
