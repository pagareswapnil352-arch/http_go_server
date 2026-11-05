package main

import (
	"log"

	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("content-type", "application/json")

		w.Write([]byte(`{"id":67,"name":"swapnil pagare","email":"pagareswapnil352@gmail.com"}`))
	})

	err := http.ListenAndServe(":8090", mux)

	if err != nil {

		log.Fatal("ListenAndServe: ", err)
	}
}
