package main

import "fmt"

func main() {
	a := -1
	b := 1
	a, b = b, a
	fmt.Println(" a: ", a)
	fmt.Println(" b: ", b)

}
