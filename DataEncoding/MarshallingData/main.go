package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`
}

func main() {
	vex := User{
		"Vex",
		20,
		"1234567890",
		true,
	}

	byteSlice, err := json.MarshalIndent(vex, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byteSlice))
}
