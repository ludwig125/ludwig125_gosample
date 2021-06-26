package main

import (
	"fmt"
	"sync"
)

// https://github.com/kat-co/concurrency-in-go-src/blob/4e55fd7f3f5b9c5efc45a841702393a1485ba206/gos-concurrency-building-blocks/the-sync-package/pool/fig-sync-pool.go

func main() {
	var mu sync.Mutex
	count := 0

	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			mu.Lock()
			defer mu.Unlock()
			count++
			fmt.Println("Creating new instance:", count)
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem // <1>
		},
	}

	fmt.Println("1")
	// Seed the pool with 4KB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	fmt.Println("2")

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()

			mem := calcPool.Get().(*[]byte) // <2>
			defer calcPool.Put(mem)

			// Assume something interesting, but quick is being done with
			// this memory.
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}
