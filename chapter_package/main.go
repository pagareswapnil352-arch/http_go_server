package main

import (
	"fmt"
	"myproject/calculator"
	"myproject/login"
)

func main() {

	result1 := calculator.Add(10, 20)

	result2 := calculator.Subtract(99, 100)

	fmt.Println("the addition is:", result1)

	fmt.Println("the addition is:", result2)

	fmt.Println(calculator.Add(100000, 200000))

	user := login.Details{

		Name:     "swapnil",
		Age:      18,
		Username: "pagareswapnil352@gmail.com",
	}

	fmt.Println("The details are:", user)
}
