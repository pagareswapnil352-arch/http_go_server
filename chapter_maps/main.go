package main

import "fmt"

func main() {

	studentgrade := make(map[string]int)
	studentgrade["rahul"] = 90
	studentgrade["ahmed"] = 80
	studentgrade["aditya"] = 70

	fmt.Println("the grade of aditya =", studentgrade["aditya"])

	fmt.Println("the grade of ahmed =", studentgrade["ahmed"])
	fmt.Println("the grade of rahul =", studentgrade["rahul"])

	grade, exist := studentgrade["rahul"]

	fmt.Println("the grade of aditya", grade)
	fmt.Println("exist", exist)
}
