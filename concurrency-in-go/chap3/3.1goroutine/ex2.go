package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	hello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}
	go hello()

	wg.Wait()
}
