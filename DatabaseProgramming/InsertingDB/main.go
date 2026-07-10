package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // To register the driver.
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(100) UNIQUE NOT NULL
);
`

func main() {
	dsn := "host=localhost port=5432 user=postgres password=admin dbname=go_db sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing the database:")
		}
	}() // can be replaced with db.Close() but this is a better way to handle errors

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database successfully!")

	//createTable(db)
	//commented because we already created the table in the previous example

	lastID, err := insertUser(db, "vex", "vex-15@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Last user ID is ", lastID)
}

func createTable(db *sql.DB) {
	res, err := db.Exec(schema)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Schema created successfully!", res)
}

func insertUser(db *sql.DB, name string, email string) (int, error) {
	var id int
	stmt := `INSERT INTO users(name, email) VALUES($1,$2) RETURNING id`

	err := db.QueryRow(stmt, name, email).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
