package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // To register the driver.
)

func main() {
	dsn := "host=localhost port=5432 user=postgres password=admin dbname=go_db sslmode=disable"
	_ = os.Remove(dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing the database:")
		}
	}()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database successfully!")
}
