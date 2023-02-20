package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var justString string

func createHugeString(n int, f *os.File) {
	var b strings.Builder
	b.Grow(1048577)
	for i := 0; i < n; i++ {
		var s string = strconv.Itoa(i)
		f.WriteString(s)
	}

}

func someFunc() {
	f, err := os.Create("ourstring.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	createHugeString(1<<10, f)
	file, err := os.Open("ourstring.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	var i int = 0
	var s string = ""
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
		k := int(math.Min(100, float64(len(string(data[:n])))))
		s += string(data[:n])[:k]
		i += k

		if i >= 100 {
			break
		}
	}
	fmt.Print(s)
}

func main() {
	someFunc()

}
