package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.AllowedDomains = []string{"www.maxiquim.com.br"}
	c.AllowURLRevisit = false
	// c.Async = true

	c.OnHTML("html body", func(e *colly.HTMLElement) {
		// e.DOM.Find("script").Remove()
		// fmt.Println(e.DOM.Html())
		fmt.Println("Got the page html for:", e.Request.URL)
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// fmt.Printf("Found a new link: %s\n", link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.maxiquim.com.br/")

	// c.Wait()
}
