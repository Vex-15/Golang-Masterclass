package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Phone    string  `json:"phone"`
	IsActive bool    `json:"is_active"`
	Profile  profile `json:"profile"`
}

type profile struct {
	Url string `json:"url"`
}

var payload = `{
  "name": "Vex",
  "age": 20,
  "phone": "1234567890",
  "is_active": true
	"profile": {
		"url": "https://github.com/Vex-15"
	}
}
`

func main() {
	var u User

	err := json.Unmarshal([]byte(payload), &u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}
