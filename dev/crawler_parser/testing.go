package main

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
	"golang.org/x/net/html"
)

func ParseHTML(pageHTML string) (pageText string, err error) {
	textTags := []string{
		"a",
		"p", "span", "em", "string", "blockquote", "q", "cite",
		"h1", "h2", "h3", "h4", "h5", "h6",
	}

	tag := ""
	enter := false

	var text string

	r := strings.NewReader(pageHTML)
	tokenizer := html.NewTokenizer(r)
	for {
		tt := tokenizer.Next()
		token := tokenizer.Token()

		err := tokenizer.Err()
		if err == io.EOF {
			break
		}

		switch tt {
		case html.ErrorToken:
			log.Fatal(err)
		case html.StartTagToken, html.SelfClosingTagToken:
			enter = false

			tag = token.Data
			for _, ttt := range textTags {
				if tag == ttt {
					enter = true
					break
				}
			}
		case html.TextToken:
			if enter {
				data := strings.TrimSpace(token.Data)

				if len(data) > 0 {
					text += data + "\n"
				}
			}
		}
	}

	return text, nil
}

func CrawlURL(resultUrl string) {
	// Parse URL
	u, err := url.Parse(resultUrl)
	if err != nil {
		log.Fatal(err)
	}
	domain := u.Hostname()
	fmt.Println("Domain: ", domain)

	c := colly.NewCollector()
	c.AllowedDomains = []string{domain}
	c.AllowURLRevisit = false
	// c.Async = true

	c.OnHTML("html body", func(e *colly.HTMLElement) {
		pageHTML, _ := e.DOM.Html()
		// fmt.Println(reflect.TypeOf(pageHtml))
		pageText, err := ParseHTML(pageHTML)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(pageText)
		// pageText, err := ParseHTML(e.Response.Request.Body)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(pageText)
		// pageText, err := ParseHTML()
		// fmt.Println(e.DOM.Html())
		// fmt.Println("Got the page html for:", e.Request.URL)

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

	c.Visit(resultUrl)

	// c.Wait()
}

func main() {
	requestUrl := "https://www.maxiquim.com.br/"
	CrawlURL(requestUrl)
}
