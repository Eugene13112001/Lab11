package main

import (
	"fmt"
)

func BinarySearch(a []int, x int) int {
	r := -1
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == x {
			r = mid
			break
		} else if a[mid] < x {
			start = mid + 1
		} else if a[mid] > x {
			end = mid - 1
		}
	}
	return r
}

func main() {

	s := 2
	x := []int{1, 2, 3, 4, 5}
	i := BinarySearch(x, s)
	fmt.Printf("BinarySearch: found  %d\n", i+1)
}
