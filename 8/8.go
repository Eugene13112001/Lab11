package main

import (
	"fmt"
	"unsafe"
)

func On(n int64, i int) int64 {
	return n | 1<<(i)
}

func Off(n int64, i int) int64 {
	return n & ^(1 << (i))
}

func main() {

	on := On(128, 5)
	off := Off(3, 1)

	fmt.Printf("%b\n", *(*uint64)(unsafe.Pointer(&on)))
	fmt.Printf("%b\n", *(*uint64)(unsafe.Pointer(&off)))

}
