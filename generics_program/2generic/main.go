package main

import "fmt"

func print[T int | bool | string](items []T) {
	for index, value := range items {
		fmt.Println(index, value)
	}
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func add[T Number](num1, num2 T) T {
	return num1 * num2
}

func main() {
	print([]int{2, 3, 4, 5, 5, 6})
	print([]string{"swapnil", "aditya", "archit"})
	print([]bool{true, false, true, true})

	fmt.Println(add(10, 20))
}
