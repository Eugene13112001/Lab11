package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan int)

	osSignals := make(chan os.Signal, 1)

	signal.Notify(osSignals, syscall.SIGINT)

	n := 10

	for i := 0; i < n; i++ {
		go func() {
			for v := range c {
				fmt.Println("Number: ", v)
			}
		}()
	}

arr:
	for {
		select {
		case <-osSignals:
			close(c)
			fmt.Println("end")
			break arr
		default:
			c <- rand.Int()
		}
	}
}
