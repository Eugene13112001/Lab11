package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	c := make(chan int)
	wg.Add(1)
	go func() {
		fmt.Println("routine started")

		select {
		case <-c:
			fmt.Println("routine stopped")
			wg.Done()
			return

		}

	}()

	c <- 1

	close(c)
	wg.Wait()

	// context
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		fmt.Println("routine started")
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("routine stopped")
				wg.Done()
				return
			default:
				i++

			}

		}
	}(ctx)

	cancel()
	wg.Wait()

	var bb int64 = 0
	wg.Add(1)
	go func(a *int) {
		fmt.Println("last routine started")
		for *a < 1 {
			time.Sleep(100 * time.Second)
			fmt.Printf("last routine running\n")
		}
		fmt.Println("last routine stopped")
		wg.Done()
	}(&bb)

	bb += 1
	wg.Wait()

}
