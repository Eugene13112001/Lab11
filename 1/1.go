package main

import "fmt"

type Human struct {
	Name string
}

func (h Human) SayName() {
	fmt.Println("My name is " + h.Name)
}

type Action struct {
	Human
}

func main() {
	h := Human{
		Name: "Eugene",
	}

	a := Action{
		Human: Human{
			Name: "Eugene",
		},
	}

	h.SayName()

	a.SayName()

}
