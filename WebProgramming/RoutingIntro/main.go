package main

import (
	"log"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Welcome to my Website! ~Vex-15</h1>"))
}

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	// 	w.Header().Set("Content-Type","text/html; charset=utf-8")
	// 	htmlContent := `
	// 	<!DOCTYPE html>
	// 	<html>
	// 	<head><title></title></head>
	// 	<body>
	// 	<h1>Welcome to my website!</h1>
	// 	<p>This is a simple web page served by Go.  ~Vex-15</p>
	// 	</body>
	// 	</html>
	// 	`
	// 	w.Write([]byte(htmlContent))

	// })

	mux := &MyHandler{}

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}
