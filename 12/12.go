package main

import "fmt"

func main() {

	a := []string{"cat", "cat", "dog", "cat", "tree"}

	myset := make(map[string]interface{})
	for _, e := range a {
		myset[e] = struct{}{}
	}
	fmt.Println("My set:")
	for i := range myset {
		fmt.Print(i, " ")
	}
}
