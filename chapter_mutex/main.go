package main

import (
	"fmt"
	"sync"
)

type insta struct {
	views int

	comments int

	mu sync.Mutex
}

func (i *insta) print(wg *sync.WaitGroup) {

	defer wg.Done()
	i.mu.Lock()
	i.views++
	i.comments++
	i.mu.Unlock()
	fmt.Println(i.views, i.comments)
}

func main() {

	var wg sync.WaitGroup
	values := insta{views: 0, comments: 0}

	for i := 1; i <= 100; i++ {

		go values.print(&wg)

		wg.Add(1)

	}
	wg.Wait()
}
