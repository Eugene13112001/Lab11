package main

import (
	"fmt"
	"time"
)

func main() {
	second := 2

	stop := time.After(time.Duration(second) * time.Second)

	c := make(chan int)

	go func() {
		i := 0
		for {
			fmt.Printf(" write %d", i)
			c <- i
			i++
			time.Sleep(100 * time.Millisecond)

		}
	}()

	go func() {
		for v := range c {
			fmt.Printf(" read %d", v)
		}
	}()

	<-stop

}
