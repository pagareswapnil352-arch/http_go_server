package main

import "fmt"

func main() {

	// below is boolean example
	//	condition1 := true

	//	condition2 := false

	//answer := condition1 && condition2

	//	answer2 := condition1 || condition2

	//fmt.Println(answer, answer2)

	var score int = 35

	if score >= 90 {
		fmt.Println("outstanding")

	} else if score >= 75 {

		fmt.Println("excellent")

	} else {
		fmt.Println("needs improvement")
	}

}
