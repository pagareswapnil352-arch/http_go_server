package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type person struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Address string `json:"Address"`
	}

	p := person{Name: "swapnil",

		Age:     18,
		Address: "khadki",
	}
	// marshal
	data, err := json.Marshal(p)

	if err != nil {

		panic(err)
	}

	fmt.Println(string(data))

	// unmarshal

	var persondata person

	err = json.Unmarshal(data, &persondata)

	fmt.Println(persondata)
}
