package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	waitNum := 3
	wg.Add(waitNum)
	hello := func(i int) {
		defer wg.Done()
		fmt.Printf("hello %d\n", i)
	}

	for i := 0; i < waitNum; i++ {
		go hello(i)
	}

	wg.Wait()
}
