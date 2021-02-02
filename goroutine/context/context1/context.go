package main

import (
	"context"
	"log"
	"time"
)

func heavyFunc(ctx context.Context) error {
	time.Sleep(5 * time.Second)
	return nil
}

func exec(ctx context.Context) error {
	errChan := make(chan error, 1)

	go func() {
		errChan <- heavyFunc(ctx)
	}()

	select {
	case <-ctx.Done(): // キャンセルが発生した場合
		return ctx.Err()
	case err := <-errChan: // heavyFuncの結果を取得した場合
		return err
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second) // 全体2秒でタイムアウト
	defer cancel()

	if err := exec(ctx); err != nil {
		log.Fatal(err)
	} else {
		log.Print("ok")
	}
}
