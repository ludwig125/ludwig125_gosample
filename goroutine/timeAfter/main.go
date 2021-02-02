// _Timeouts_ are important for programs that connect to
// external resources or that otherwise need to bound
// execution time. Implementing timeouts in Go is easy and
// elegant thanks to channels and `select`.

package main

import (
	"context"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// ch := make(chan string, 1)
	// go func(ctx context.Context) {
	// 	defer close(ch)
	// 	max := 10000000
	// 	for i := 0; i <= max; i++ {
	// 		select {
	// 		case <-ctx.Done():
	// 			log.Printf("context cancelled. count: %d", i)
	// 			return
	// 		default:
	// 		}
	// 		//fmt.Println(i)
	// 		time.Sleep(1 * time.Nanosecond)
	// 	}
	// 	ch <- string(max)
	// }(ctx)
	ch := task(ctx)

	select {
	//case <-ch:
	case <-time.After(1 * time.Second):
		cancel()
		log.Println("timeout after 1s")
	}
	<-ch
	log.Printf("Total time: %fs", time.Since(start).Seconds())
}

func task(ctx context.Context) <-chan string {
	ch := make(chan string, 1)
	go func() {
		defer close(ch)
		i := 0
		for {
			select {
			case <-ctx.Done():
				log.Printf("context cancelled. count: %d", i)
				return
			default:
			}
			i++
		}
	}()
	return ch
}
