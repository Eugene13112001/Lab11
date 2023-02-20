package main

import (
	"fmt"
	"math"
)

func main() {
	var nums = []int{2, 4, 6, 8, 10}

	p := make(chan float64)
	for _, num := range nums {
		go func(number int) {
			pow := math.Pow(float64(number), 2)
			fmt.Println("Done : ", pow)
			p <- pow
		}(num)
	}

	sum := 0.0
	for range nums {
		sum += <-p
	}

	fmt.Println(sum)

}
