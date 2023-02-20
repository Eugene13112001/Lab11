package main

import (
	"fmt"

	"sync"
)

func main() {

	nums := make([]int, 7)
	for i := 0; i < 7; i++ {
		nums[i] = i
	}

	chan1 := make(chan int)
	chan2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for _, num := range nums {
			chan1 <- num
		}
		close(chan1)
		wg.Done()
	}()

	go func() {
		for num := range chan1 {
			chan2 <- num * num
		}
		close(chan2)
		wg.Done()
	}()

	go func() {
		for num := range chan2 {
			fmt.Printf("%d\n", num)
		}
		wg.Done()
	}()

	wg.Wait()

}
