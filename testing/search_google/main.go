package main

import (
	"context"
	"fmt"

	googlesearch "github.com/rocketlaunchr/google-search"
)

var ctx = context.Background()

func main() {

	q := "Facebook"

	opts := googlesearch.SearchOptions{
		CountryCode:  "br",
		LanguageCode: "pt-br",
		Limit:        5,
	}

	returnLinks, err := googlesearch.Search(ctx, q, opts)
	if err != nil {
		fmt.Printf("something went wrong: %v", err)
		return
	}

	if len(returnLinks) == 0 {
		fmt.Printf("no results returned: %v", returnLinks)
	}

	fmt.Println("Results are: ", returnLinks)
}
