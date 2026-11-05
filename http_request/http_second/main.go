package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://example.com"

	response, err := http.Get(url)

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
