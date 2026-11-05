package main

import "fmt"

func main() {

	names := []string{"swapnil", "aditya", "archit"}

	for index, name := range names {

		fmt.Println(index, name)
	}

}
