package main

import (
	"fmt"
	"unicode"
)

func CheckUnique(s string) bool {

	smbl := make(map[rune]int)
	for _, r := range s {
		r = unicode.ToLower(r)
		if _, ok := smbl[r]; ok {
			return false
		}
		smbl[r] = 1
	}
	return true
}

func main() {

	str := []string{"abcd", "abCdefAafhh", "aabcdvfyu", "aA", ""}
	for _, s := range str {
		fmt.Printf("%s - %v\n", s, CheckUnique(s))
	}

}
