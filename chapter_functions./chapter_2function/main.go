package main

import "fmt"

func swapnil() {

	fmt.Println("hii swapnil")

}
func add(sum1 int, sum2 int) (int, int) {

	//result := sum1 + sum2
	//result2 := sum1 * sum2

	//return result, result2
	return sum1 + sum2, sum2 * sum1
}

func main() {

	swapnil()
	//	answer2, answer := add(20, 30)
	//	fmt.Println(answer, answer2)
	fmt.Println(add(50, 20))
}
