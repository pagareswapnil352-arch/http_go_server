package main

import (
	"fmt"
	"net/url"
)

func main() {

	var rawurl = "https://example.com"

	u, err := url.Parse(rawurl)

	if err != nil {

		panic(err)

	}

	fmt.Println(u.Path)
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)

}
