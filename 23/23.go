package main

import (
	"fmt"
)

func RemoveElement(s []string, i int) []string {

	return append(s[:i], s[i+1:]...)
}

func main() {

	users := []string{"Eugene", "Sun", "Yellow", "Sam"}

	ns := RemoveElement(users, 2)
	fmt.Printf("%v\n", ns)

}
