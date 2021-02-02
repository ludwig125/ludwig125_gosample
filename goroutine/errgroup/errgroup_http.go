package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// メッセージの内容を出力するタスク
	messages := []string{"first", "second", "third"}
	var g errgroup.Group
	for _, message := range messages {
		message := message
		g.Go(func() error {
			time.Sleep(2 * time.Second)
			if message == "three" {
				return errors.New("error")
			}
			fmt.Printf("message=%s\n", message)
			return nil
		})
	}
	// タスクの完了を待合せる
	// いずれかのタスクでエラーになると、err != nil になる
	if err := g.Wait(); err == nil {
		fmt.Printf("tasks finished successfully\n")
	} else {
		fmt.Printf("some task failed\n")
	}
}
