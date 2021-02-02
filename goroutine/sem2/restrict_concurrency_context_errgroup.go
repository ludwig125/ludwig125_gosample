package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := doTask(); err != nil {
		log.Printf("error occured. %v", err)
	}
	log.Println("finished")
}

const concurrency = 2 // 最大同時並列実行数

var errFlag bool = true

func doTask() error {
	numbers := []int{1, 2, 3, 4, 5, 6}

	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sem := make(chan struct{}, concurrency)
	for _, num := range numbers {
		sem <- struct{}{} // チャネルに送信
		log.Printf("num: %d", num)

		n := num
		eg.Go(func() error {
			defer func() { <-sem }()

			log.Printf("goroutine num: %d", n)
			select {
			case <-ctx.Done():
				//return ctx.Err()
				return nil
			default:
			}
			if err := fnA(n); err != nil {
				// エラーが発生したら他の処理はキャンセル
				cancel()
				time.Sleep(1 * time.Second) // 処理をわかりやすくするため
				return fmt.Errorf("failed to A, %v", err)
			}
			time.Sleep(1 * time.Second) // 処理をわかりやすくするため
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		defer close(sem)
		cancel()
		return err
	}
	return nil
}

func fnA(n int) error {
	log.Println("do fnA.")
	if errFlag {
		log.Printf("--> failed to do fnA. num: %d", n)
		return fmt.Errorf("error A. num: %d", n)
	}
	log.Printf("--> succeeded to do fnA. num: %d", n)
	return nil
}
