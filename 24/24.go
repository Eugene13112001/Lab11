package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p *Point) GetY() float64 {
	return p.y
}

func (p *Point) GetX() float64 {
	return p.x
}

func CountDist(a, b *Point) float64 {
	return math.Sqrt(math.Sqrt(a.GetX()-b.GetX()) + math.Sqrt(a.GetY()-b.GetY()))
}

func main() {
	a := NewPoint(-1.23, 6.74)
	b := NewPoint(10.2, 3.40)

	fmt.Println("Расстояние ", CountDist(&a, &b))

}
