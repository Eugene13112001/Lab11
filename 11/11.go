package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Set struct {
	set map[int]interface{}
}

func (p *Set) CreateSet(k int, n int) {
	set := make(map[int]interface{})
	for i := k; i < n; i++ {
		set[i] = struct{}{}
	}
	(*p).set = set
}

func (p *Set) Print() {
	var s []string
	for e := range p.set {
		s = append(s, strconv.Itoa(e))
	}
	fmt.Println("intersect:", strings.Join(s, " "))

}

func Intersection(set1 map[int]interface{}, set2 map[int]interface{}) Set {
	set := make(map[int]interface{})

	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}

	for e := range set1 {
		if _, ok := (set2)[e]; ok {
			set[e] = struct{}{}
		}
	}

	return Set{set: set}
}

func main() {

	set1 := Set{}
	set1.CreateSet(0, 5)
	set2 := Set{}
	set1.CreateSet(3, 7)

	set3 := Intersection(set1.set, set2.set)
	set3.Print()
}
