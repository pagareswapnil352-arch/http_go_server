package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Create("example.txt")

	if err != nil {

		fmt.Println("someting went wrong")

	}
	var content = "hii i am swapnil"
	_, err = io.WriteString(file, content)

	if err != nil {

		fmt.Println("someting went wrong")

	}
	defer file.Close()

	data, err := os.ReadFile("example.txt")

	if err != nil {

		fmt.Println("someting went wrong")

	}

	fmt.Println("the data are:", string(data))
}
