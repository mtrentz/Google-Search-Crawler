package main

import (
	"fmt"
	"log"
)

type SqlError struct {
	errType     string
	description string
}

func (r *SqlError) Error() string {
	return fmt.Sprintf("SQL Error type %s, description: %s", r.errType, r.description)
}

func squareit(x int) (res int, err error) {
	if x > 10 {
		return 0, &SqlError{"add", "my string"}
	} else {
		return x * x, nil
	}
}

func main() {
	val, err := squareit(11)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)
}
