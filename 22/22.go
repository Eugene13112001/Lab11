package main

import (
	"fmt"
	"math"
	"math/big"
)

func sum(a, b *big.Int) *big.Int {
	var c big.Int
	return c.Add(a, b)
}

func dif(a, b *big.Int) *big.Int {
	var c big.Int
	return sum(a, c.Neg(b))
}

func mul(a, b *big.Int) *big.Int {
	var c big.Int
	return c.Mul(a, b)
}

func div(a, b *big.Int) *big.Int {
	var c big.Int
	return c.Quo(a, b)
}

func main() {

	a := big.NewInt(10 * int64(math.Pow(2.0, 40.00)))
	b := big.NewInt(int64(math.Pow(2.0, 22.00)))

	fmt.Printf("a = %v, b = %v\n", a, b)
	fmt.Printf("a + b = %v\n", sum(a, b))
	fmt.Printf("a - b = %v\n", dif(a, b))
	fmt.Printf("a * b = %v\n", mul(a, b))
	fmt.Printf("a / b = %v\n", div(a, b))
}
