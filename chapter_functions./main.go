package main

import "fmt"

func add(x int, y int) (int, int) {
	return x + y, x - y
}

func main() {
	y, x := add(10, 5)

	fmt.Println(x, y)
}
