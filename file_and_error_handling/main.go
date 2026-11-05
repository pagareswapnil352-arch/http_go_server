package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Create("example.txt")

	if err != nil {

		fmt.Println("there is a error:", err)

	}

	var content = "hii i am swapnil"

	_, err = io.WriteString(file, content)

	if err != nil {

		fmt.Println("error while creating a file", err)
	}
	fmt.Println("files had created")

	data, err := os.ReadFile("example.txt")

	if err != nil {

		fmt.Println("there is an error")

	}

	fmt.Println("the data is", data)
	fmt.Println(string(data))

}
