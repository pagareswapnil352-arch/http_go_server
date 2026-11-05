package main

import "fmt"

type car struct {
	name string
	year int
}

func (c car) carname() {

	fmt.Println("the car name :", c.name)
	fmt.Println("the year of manufacture", c.year)

}

func greeting() {

	fmt.Println("thank you for checking out us")
}

func main() {

	car1 := car{name: "mercedes", year: 2018}

	car1.carname()

	car2 := car{name: "bmw", year: 2020}

	car2.carname()

	greeting()

}
