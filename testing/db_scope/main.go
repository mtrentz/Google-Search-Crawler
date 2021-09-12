package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func main() {
	fmt.Println("Drivers:", sql.Drivers())
	var err error
	DB, err = sql.Open("mysql", "root:7622446@tcp(127.0.0.1:3306)/web_scrape")
	if err != nil {
		log.Fatal("Unable to open connection to db")
	}
	defer DB.Close()

	res := QueryDB("https://www.maxiquims.com.br/")
	fmt.Println(res)

}
