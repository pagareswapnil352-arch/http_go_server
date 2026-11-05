package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) introduce() {
	fmt.Println("My name is", p.name)
	fmt.Println("My age is", p.age)
}

func main() {
	person := Person{
		name: "Rahul",
		age:  20,
	}

	person.introduce()
}
