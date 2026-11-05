package main

import (
	"fmt"
	"sync"
)

func numbers(num int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("the numbers are:", num)
}

func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go numbers(i, &wg)

	}

	wg.Wait()
}
