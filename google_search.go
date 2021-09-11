package main

import (
	"context"
	"fmt"

	googlesearch "github.com/rocketlaunchr/google-search"
)

var ctx = context.Background()

func GoogleSearch(query string) ([]googlesearch.Result, error) {
	opts := googlesearch.SearchOptions{
		CountryCode:  "br",
		LanguageCode: "pt-br",
	}

	returnLinks, err := googlesearch.Search(ctx, query, opts)
	if err != nil {
		fmt.Printf("something went wrong: %v", err)
		return returnLinks, err
	}

	return returnLinks, err
}
