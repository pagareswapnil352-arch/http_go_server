package main

import "fmt"

func main() {

	marks := [5]int{30, 40, 50, 60, 70}

	marks[0] = 100
	fmt.Println(marks)
	fmt.Println(len(marks))

	var num [5]int

	num[0] = 10
	num[1] = 20
	num[2] = 30
	num[3] = 40
	num[4] = 50

	fmt.Println(num)
}
