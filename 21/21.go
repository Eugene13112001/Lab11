package main

import (
	"fmt"
)

type Transport interface {
	drive()
}

type Ship interface {
	move()
}

type Man struct {
	Name string
}

type Bike struct {
}

type Liner struct {
}

type Adapter struct {
	Ship
}

func (c Bike) drive() {
	fmt.Println("Автомобиль едет")
}

func (c Liner) move() {
	fmt.Println("Корабль плывет")
}

func (c Adapter) drive() {
	c.move()
}

func (c Man) drive(tran Transport) {
	tran.drive()
}

func main() {

	man := Man{Name: "Max"}
	var bike Transport = Bike{}
	man.drive(bike)
	var liner Ship = Liner{}
	var adapt Adapter = Adapter{Ship: liner}
	man.drive(adapt)

}
