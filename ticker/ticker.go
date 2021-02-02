package main

import (
	"context"
	"fmt"
	"time"
)

func ticker(ctx context.Context) error {
	t := time.NewTicker(1 * time.Second) //1秒周期の ticker
	defer t.Stop()

	for {
		select {
		case now := <-t.C:
			fmt.Println(now.Format(time.RFC3339))
		case <-ctx.Done():
			fmt.Println("Stop child")
			return ctx.Err()
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ticker(ctx)
}
