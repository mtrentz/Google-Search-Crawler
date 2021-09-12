package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Drivers:", sql.Drivers())
	db, err := sql.Open("mysql", "root:7622446@tcp(127.0.0.1:3306)/web_scrape")
	if err != nil {
		log.Fatal("Unable to open connection to db")
	}
	defer db.Close()

	// // SELECT
	// val := "maxiquim"

	// var id int64

	// // CHECK IF ALREADY EXISTS
	// // err = db.QueryRow("SELECT id FROM queries WHERE query_text=?", val).Scan(&id)
	// err = db.QueryRow("SELECT id FROM queries WHERE query_text=?", val).Scan(&id)
	// if err != nil && err != sql.ErrNoRows {
	// 	// TODO: Return error and log
	// 	panic(err.Error()) //
	// }
	// // If different from 0, query was already added
	// if id != 0 {
	// 	fmt.Println("Already exists, ID: ", id)
	// 	os.Exit(1)
	// }
	// fmt.Println("Not yet in db, continuing...")

	// CHECK IF ALREADY EXISTS, if so, throws error because I don't want to scrape that domain again
	var exists bool
	testUrl := "https://www.maxiquims.com.br/"
	err = db.QueryRow("SELECT EXISTS (SELECT * FROM query_results WHERE url=?)", testUrl).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		// TODO: Return error here
		log.Fatal("test")
	}
	if exists {
		fmt.Println("Exists!")
	} else {
		fmt.Println("Doesnt exists")
	}

}
