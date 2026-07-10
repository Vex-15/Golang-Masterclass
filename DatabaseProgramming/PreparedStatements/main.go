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

type User struct {
	ID    int
	Name  string
	Email string
}

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

	// lastID, err := insertUser(db, "vex", "vex-15@gmail.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Last user ID is ", lastID)
	//commented because we already inserted data in the previous example

	// user, err := fetchUserById(db, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Fetched user", user)

	// users, err := fetchAllUsers(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Fetched all users", users)
	//commented as we fetched the data in the previous example

	_, err = insertUserWithPrepare(db, "abc", "abc@example.com")
	if err != nil {
		log.Fatal(err)
	}
	_, err = insertUserWithPrepare(db, "pqr", "pqr@example.com")
	if err != nil {
		log.Fatal(err)
	}
	_, err = insertUserWithPrepare(db, "xyz", "xyz@example.com")
	if err != nil {
		log.Fatal(err)
	}

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

func insertUserWithPrepare(db *sql.DB, name string, email string) (int, error) {
	var id int
	stmt, err := db.Prepare(`INSERT INTO users(name, email) VALUES($1,$2) RETURNING id`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	defer fmt.Println("Inserted user with prepared statement successfully!")

	err = stmt.QueryRow(name, email).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func fetchUserById(db *sql.DB, id int) (*User, error) {
	stmt := `SELECT id,name, email FROM users WHERE id=$1`
	user := &User{}
	err := db.QueryRow(stmt, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func fetchAllUsers(db *sql.DB) ([]User, error) {
	stmt := `SELECT id,name, email FROM users`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
