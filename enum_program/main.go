package main

import "fmt"

type value string

const (
	choice1 value = "dost"
	choice2 value = "bro"
	choice3 value = "bhai"
)

func printchoice(p value) {

	fmt.Println("the value is:", p)
}

func main() {

	printchoice(choice1)

	fmt.Println(choice1, choice2, choice3)

}
