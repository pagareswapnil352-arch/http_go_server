package main

import "fmt"

type car struct {
	brand string

	year int
}

func main() {

	ferrari := car{

		brand: "ferrari",

		year: 2020}

	fmt.Println(ferrari.brand, ferrari.year)

}
