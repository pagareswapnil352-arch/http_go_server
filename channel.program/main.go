package main

import (
	"fmt"
)

func channel(message chan int) {

	for value := range message {

		fmt.Println("the message is ", <-message, value)
	}
}

func main() {

	message := make(chan int)

	go channel(message)

	for i := 1; i <= 100; i++ {

		message <- i

	}
}
