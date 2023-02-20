package main

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
// 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
// градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
//

import (
	"fmt"
	"math"
	"sort"
)

type Temps struct {
	numbers []float64
	minT    float64
	maxT    float64
}

func main() {
	a := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -7.8, 3.5, 0.0}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	b := []Temps{}
	c := -1

	for _, v := range a {
		if c == -1 {
			c += 1
			b = append(b, Temps{numbers: []float64{v}, minT: v, maxT: v})
		} else {
			if (math.Abs(v-float64(b[c].minT)) < 10) && (math.Abs(v-float64(b[c].maxT)) < 10) {
				b[c].minT = math.Min((b[c].minT), v)
				b[c].maxT = math.Max((b[c].maxT), v)
				b[c].numbers = append(b[c].numbers, v)
			} else {
				c += 1
				b = append(b, Temps{numbers: []float64{v}, minT: v, maxT: v})
			}
		}

	}
	for _, v1 := range b {
		fmt.Println(v1.numbers)
		fmt.Print("\n")

	}

}
