package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"golang.org/x/net/context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		select {
		case <-signalCh:
			cancel()
			return
		}
	}()

	<-doTask(ctx)
	// for v := range ch {
	// 	fmt.Println(v)
	// }

}

func doTask(ctx context.Context) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("received done")
				// Do something before terminated
				time.Sleep(500 * time.Millisecond)
				return
			default:
			}
			// do something
			fmt.Println("sleep 1. count:", i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}
