package main

import (
	"fmt"
	"math/rand"
)

func main() {
	newRandStream := func(done <-chan interface{}) chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)

}
