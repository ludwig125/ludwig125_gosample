package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	doConcurrency()
	fmt.Println("finished")
}

const concurrency = 2

func doConcurrency() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup
	for _, num := range numbers {
		sem <- struct{}{}

		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			//defer func() { <-sem }()
			defer func() {
				fmt.Println("---")
				time.Sleep(1 * time.Second)
				<-sem
			}()
			fnA(n)

		}(num)
	}

	wg.Wait()
}

func fnA(n int) {
	fmt.Printf("do fnA. num: %d \n", n)
}
