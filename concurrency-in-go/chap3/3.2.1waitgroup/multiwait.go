// https://github.com/kat-co/concurrency-in-go-src/blob/4e55fd7f3f5b9c5efc45a841702393a1485ba206/gos-concurrency-building-blocks/the-sync-package/waitgroup/fig-bulk-add.go

package main

import (
	"fmt"
	"sync"
)

func main() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
