package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mutex sync.Mutex
	count int
}

func main() {

	c := Counter{count: 0}

	c1 := 0

	var wg sync.WaitGroup
	n := 100
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(c2 *int) {

			*c2 += 1
			c.mutex.Lock()
			c.count++
			c.mutex.Unlock()

			wg.Done()
		}(&c1)
	}
	wg.Wait()
	fmt.Println("main", c1)
	fmt.Println("mutex ", c.count)
}
