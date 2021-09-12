package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Set up log file
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}
	log.SetOutput(logFile)

	test := "Testando 123"
	log.Printf("Testando, %s", test)
	fmt.Println("Logged inside main!")

	res := MyFunc()
	fmt.Println(res)
}
