package main

import (
	"net/http"

	"fmt"
)

func main() {
	// 1. Setup the route

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			fmt.Fprintln(w, "your information is this")

		} else {
			fmt.Fprintln(w, "submit your information ")
		}

	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "I am a go developer")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Contact me at example@email.com")
	})

	err := http.ListenAndServe(":8081", nil)
	fmt.Println("Server stopped:", err)

}
