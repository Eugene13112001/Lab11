package main

import (
	"fmt"
)

func RevWord(s string) string {
	w := s
	n := len(w)
	var str string = ""

	for i := (n - 1); i >= 0; i-- {
		str += string(w[i])
	}
	return str
}

func main() {
	s := "He came to office"
	fmt.Println("NewWord:", RevWord(s))
}
