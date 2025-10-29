package main

import (
	"fmt"

	"github.com/sumit40000/book-management-api/internal/config"
)

func main() {
	fmt.Println("welcome to the book management api")

	// connecting the database
	config.ConnectDb()
}