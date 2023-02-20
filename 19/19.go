package main

import (
	"fmt"
	"os"
)

func ReverseString(s string) string {
	w := s
	n := len(w)
	var str string = ""

	for i := (n - 1); i >= 0; i-- {
		str += string(w[i])
	}
	return str
}

func main() {
	var str string

	fmt.Fscan(os.Stdin, &str)
	fmt.Println("Reversed:", ReverseString(str))

}
