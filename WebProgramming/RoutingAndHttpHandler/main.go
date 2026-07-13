package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)

	log.Println("Listening on localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Print(err)
	}

}
