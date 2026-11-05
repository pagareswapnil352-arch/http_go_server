package main

import "fmt"

func main() {

	cars := []string{"ferrari", "bmw", "audi", "mercedes"}

	fmt.Println("the following are the popular cars in worldwide :", cars)

	for index, value := range cars {

		fmt.Println(index, value)
	}

}
