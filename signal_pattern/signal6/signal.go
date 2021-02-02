package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"
)

func main() {
	res, err := doTask(newCtx())
	for v := range res {
		fmt.Println("done successfully.", v)
	}
	for e := range err {
		fmt.Printf("failed to doTask: %v", e)
		return
	}
}

func newCtx() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sCh := make(chan os.Signal, 1)
		signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM)
		<-sCh
		fmt.Println("Got signal!")
		cancel()
	}()
	return ctx
}

func doTask(ctx context.Context) (<-chan string, <-chan error) {
	resCh := make(chan string)
	errCh := make(chan error, 5)
	go func() {
		defer fmt.Println("done doTask")
		defer close(resCh)
		defer close(errCh)
		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("received done")
				// Do something before terminated
				time.Sleep(500 * time.Millisecond)
				errCh <- ctx.Err()
				return
			default:
			}

			// do something
			fmt.Println("sleep 1. count:", i)
			time.Sleep(time.Second)
		}
		resCh <- fmt.Sprintf("something")
	}()
	return resCh, errCh
}
