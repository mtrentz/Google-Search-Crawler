package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {

	file, err := os.Open("urls.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		myUrl := scanner.Text()
		u, err := url.Parse(myUrl)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("URL: %s -> Domain: %s\n", myUrl, u.Hostname())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
