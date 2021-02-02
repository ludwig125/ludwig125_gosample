package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	waitNum := 3
	wg.Add(waitNum)
	hello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}

	for i := 0; i < waitNum; i++ {
		go hello()
	}

	wg.Wait()
}
