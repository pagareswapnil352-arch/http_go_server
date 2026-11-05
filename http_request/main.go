package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("http web request")

	response, err := http.Get("https://lcu.dev")

	if err != nil {

		panic(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {

		panic(err)
	}

	fmt.Println(string(body))

}
