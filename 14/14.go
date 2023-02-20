package main

import (
	"fmt"
)

func KnowType(x interface{}) string {
	switch x.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	}
	return "unknown type"
}

func main() {

	arr := []interface{}{900, "string", false}

	for _, x := range arr {
		fmt.Println("Type: ", KnowType(x))
	}

}
