package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`
}

func main() {
	// u := User{
	// 	Name:     "Vex",
	// 	Age:      20,
	// 	Phone:    "1234567890",
	// 	IsActive: true,
	// }

	// enc = json.NewEncoder(os.Stdout)
	// if err := enc.Encode(&u); err != nil {
	// 	log.Fatal(err)
	// }
	//Above one is for encoding data to json and writing it to stdout

	//Below one is for decoding data from json and writing it to stdout
	var payload = `{
	"name": "Vex",
	"age": 20,
	"phone": "1234567890",	
	"is_active": true
	}`

	var u User
	enc := json.NewDecoder(strings.NewReader(payload))
	if err := enc.Decode(&u); err != nil {
		log.Fatal(err)
	}

	fmt.Println(u)
}
