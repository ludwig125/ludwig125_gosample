package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := send(); err != nil {
		fmt.Printf("error occured. %v\n", err)
	}
	fmt.Println("finished")
}

const concurrency = 3

var count int

func send() error {
	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sem := make(chan struct{}, concurrency)
	for num := range numbers {
		n := num
		eg.Go(func() error {
			sem <- struct{}{}
			defer func() { <-sem }()
			select {
			case <-ctx.Done():
				return nil
			default:
			}
			if err := fnA(n); err != nil {
				return fmt.Errorf("failed to A, %v", err)
			}
			fmt.Println("succeeded to do A", n)

			if err := fnB(n); err != nil {
				return fmt.Errorf("failed to B, %v", err)
			}
			fmt.Println("succeeded to do B", n)
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		cancel()
		return err
	}
	return nil
}

func fnA(n int) error {
	fmt.Printf("do fnA. count: %d\n", count)
	if count >= 3 {
		return fmt.Errorf("error A %d", n)
	}
	return nil
}
func fnB(n int) error {
	fmt.Printf("do fnB. count: %d\n", count)
	if count >= 3 {
		return fmt.Errorf("error B %d", n)
	}
	count++
	return nil
}
