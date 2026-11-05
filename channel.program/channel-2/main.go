package main

import "fmt"

func add(sum chan int, num1 int, num2 int) {

	sum <- num1 * num2

}

func main() {

	sum := make(chan int)

	go add(sum, 2, 3)

	res := <-sum

	fmt.Println(res)

	channel := make(chan string, 2)

	channel <- "swapnil"

	x := <-channel

	fmt.Println("the channel value:", x)

}
