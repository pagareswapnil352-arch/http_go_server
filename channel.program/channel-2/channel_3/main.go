package main

import (
	"fmt"
	"time"
)

func truth(done chan bool) {

	fmt.Println("the channel is", <-done)
}

func main() {

	message := make(chan int, 2)
	message <- 100

	message <- 2000
	fmt.Println(<-message, <-message)

	time.Sleep(time.Second)

	done := make(chan bool)

	go truth(done)

	done <- true

}
