package main

import (
	"fmt"
	"github.com/Ja7ad/golang-tips/orm/db"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func init() {
	// Initial ENV
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Initial DBMS
	switch os.Getenv("DBMS") {
	case "mysql":
		if err := db.MysqlConnector(); err != nil {
			panic(err)
		}
	case "sqlite":
		if err := db.SqliteConnector(); err != nil {
			panic(err)
		}
	default:
		panic("DBMS not found, please set mysql or sqlite")
	}
}

func main() {
	book := db.Books{}

	// Add New Book
	err := book.AddBook(map[string]interface{}{
		"BookName":     "Book 1",
		"author_name":  "Javad",
		"author_email": "ja7adr@gmail.com",
		"PublishDate":  time.Now(),
	})
	if err != nil {
		panic(err)
	}

	// Get list all book
	allBooks, err := book.GetAllBook()
	if err != nil {
		panic(err)
	}
	fmt.Println(allBooks)

	// Get one book by id
	oneBook, err := book.GetBook(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(oneBook)

	// Update one book by id
	if err := book.UpdateBook(1, map[string]interface{}{
		"BookName":     "Book 2",
		"author_name":  "Test Author",
		"author_email": "test@gmail.com",
		"PublishDate":  time.Now(),
	}); err != nil {
		panic(err)
	}

	// Delete one book by id
	if err := book.DeleteBook(1); err != nil {
		panic(err)
	}
}
