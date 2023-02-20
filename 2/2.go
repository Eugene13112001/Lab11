package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var nums = []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	for _, num := range nums {

		wg.Add(1)

		go func(num int) {
			pow := math.Pow(float64(num), 2)

			fmt.Println("Done : ", pow)
			wg.Done()
		}(num)
	}

	wg.Wait()

}
