package main

import "fmt"

func main() {

	// declare slice var marks []int
	// slices are the same as arrays but they are dynamic in size

	// marks:=[in]t{1,2,3,4,5}
	// marks[0] = 10
	//marks = append(marks,6,7,8,9,10)
	//fmt.println(marks)
	//marks : = make([int,5)

	marks := []int{10, 50, 30, 40, 20}
	marks = append(marks, 60, 70, 80, 90, 100)

	marks1 := make([]int, 5)

	fmt.Println(len(marks))

	marks1[0] = 10
	marks1[1] = 20
	marks1[2] = 30
	marks1[3] = 40
	marks[4] = 50

	fmt.Println(marks1)

	fmt.Println("the marks are:", marks)

}
