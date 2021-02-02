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

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	defer func() {
		// シグナルの受付を終了する
		signal.Stop(sigs)
		cancel()
	}()
	go func() {
		select {
		case sig := <-sigs: // シグナルを受け取ったらここに入る
			fmt.Println("Got signal!", sig)
			cancel() // cancelを呼び出して全ての処理を終了させる
			return
		}
	}()

	res, err := doTask(ctx)
	for v := range res {
		fmt.Println("done successfully.", v)
	}
	for e := range err {
		fmt.Printf("failed to doTask: %v", e)
		cancel() // 何らかのエラーが発生した場合、他の処理も全てcancelさせる
		return
	}
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
			// // エラー時の挙動が見たい場合はここのコメントアウトを外す
			// if i == 3 {
			// 	errCh <- fmt.Errorf("error happened")
			// 	return
			// }

			// do something
			fmt.Println("sleep 1. count:", i)
			time.Sleep(time.Second)
		}
		resCh <- fmt.Sprintf("something")
	}()
	return resCh, errCh
}
