package main

import "fmt"

func main() {

	a := 10

	p := &a
	fmt.Println(p) // p used to print the address of a

	fmt.Println(*p) // *p used to print the value of a

	*p = 20 // *p used to change the value of a
	fmt.Println(a)
}
		