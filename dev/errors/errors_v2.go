package main

import (
	"errors"
	"fmt"
)

var ErrSqlStatement = errors.New("sql: no rows in result set")
var errTest error

func squareit(x int) (res int, err error) {
	if x > 10 {
		return 0, ErrSqlStatement
	} else if x < -1 {
		fmt.Println("Here")
		return 0, errTest
	} else {
		return x * x, nil
	}
}

func main() {
	_, err := squareit(12)
	fmt.Printf("%T\n, %v\n", err, err)
	// if err == ErrSqlStatement {
	// 	fmt.Println("Sql Error")
	// } else if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("No errors")
	// }
}
