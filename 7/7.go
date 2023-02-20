package main

import (
	"fmt"
	"sync"
)

func main() {

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = i
	}

	var counter = struct {
		p int
		sync.RWMutex
		m map[int]int
	}{m: make(map[int]int), p: 0}
	var wg sync.WaitGroup
	wg.Add(len(nums))

	for _, v := range nums {
		go func(i int) {
			counter.RLock()

			counter.p += 1
			counter.m[i] = counter.p
			counter.RUnlock()

			wg.Done()
		}(v)
	}
	wg.Wait()
	for _, k := range counter.m {
		fmt.Println("Key:", k, "Value:", counter.m[k])
	}

}
