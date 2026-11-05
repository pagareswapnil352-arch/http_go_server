package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("HTTP server starting on http://localhost:8080")

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/missing", missingHandler)

	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home page - 200 OK")
}

func missingHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
