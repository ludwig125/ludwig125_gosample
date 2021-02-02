package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := task(ctx)
	select {
	case <-res:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		log.Println("timeout after 1s")
		cancel()
		log.Println(<-res)
	}
	log.Printf("Total time: %fs", time.Since(start).Seconds())
}

func task(ctx context.Context) <-chan string {
	res := make(chan string)
	go func() {
		defer close(res)
		for i := 0; i < 1000000; i++ {
			select {
			case <-ctx.Done():
				log.Printf("context cancelled. count: %d", i)
				res <- fmt.Sprintf("error: %s", ctx.Err())
				return
			default:
			}
			fmt.Println(i)
			time.Sleep(1 * time.Nanosecond)
		}
		res <- "done!"
	}()
	return res
}
