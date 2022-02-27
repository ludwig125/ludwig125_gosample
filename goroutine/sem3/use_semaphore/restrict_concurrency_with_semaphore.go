package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	start := time.Now()

	fmt.Println("runtime.GOMAXPROCS(0)", runtime.GOMAXPROCS(0))
	concurrency := 1
	if err := doTask(concurrency); err != nil {
		log.Printf("error occured. %v", err)
	}
	log.Println("finished", time.Since(start).Milliseconds())
}

type Response struct {
	res int
	err error
}

func doTask(concurrency int) error {
	ctx := context.TODO()
	numbers := []int{1, 2, 3, 4, 5, 6}

	// maxWorkers := runtime.GOMAXPROCS(0)
	// sem := semaphore.NewWeighted(int64(maxWorkers))
	// https://pkg.go.dev/golang.org/x/sync/semaphore#example-package-WorkerPool
	sem := semaphore.NewWeighted(10)

	res := make(chan Response, len(numbers))
	for _, num := range numbers {
		log.Printf("num: %d", num)

		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}

		go func(num int) {
			defer sem.Release(1)
			r, err := tenTimes(num)
			res <- Response{
				res: r,
				err: err,
			}
		}(num)
	}

	for range numbers {
		r := <-res
		if r.err != nil {
			log.Printf("failed: %v", r.err)
			continue
		}
		log.Println("res", r.res)
	}

	return nil
}

var tenTimes = func(n int) (int, error) {
	time.Sleep(1000 * time.Millisecond)
	if n%3 == 0 {
		// 3の倍数だけエラーを返す
		return 0, fmt.Errorf("error at %d", n)
	}
	return n * 10, nil
}
