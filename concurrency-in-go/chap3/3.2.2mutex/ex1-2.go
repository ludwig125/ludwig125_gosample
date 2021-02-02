package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println(time.Now())

	var count int
	//var lock sync.Mutex
	var wg sync.WaitGroup

	increment := func() {
		//	time.Sleep(500 * time.Millisecond)
		//lock.Lock()
		//defer lock.Unlock()
		count++
		fmt.Printf("Increment: %d\n", count)
	}

	decrement := func() {
		//time.Sleep(200 * time.Millisecond)
		//lock.Lock()
		//defer lock.Unlock()
		count--
		//time.Sleep(300 * time.Millisecond)
		fmt.Printf("Decrement: %d\n", count)
	}

	// インクリメント
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	// デクリメント
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
		//time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()
	fmt.Println(time.Now())
}
