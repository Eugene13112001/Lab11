package main

import (
	"fmt"
	"time"
)

func sleep(sec int) {
	<-time.After(time.Duration(sec) * time.Second)
}

func main() {
	seconds := 4
	sleep(seconds)
	fmt.Println("final")
}
