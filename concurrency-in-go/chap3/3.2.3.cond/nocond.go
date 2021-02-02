// https://github.com/kat-co/concurrency-in-go-src/blob/4e55fd7f3f5b9c5efc45a841702393a1485ba206/gos-concurrency-building-blocks/the-sync-package/cond/fig-cond-based-queue.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})    // <1>
	queue := make([]interface{}, 0, 10) // <2>

	removeFromQueue := func(delay time.Duration) {
		fmt.Println("do remove")
		time.Sleep(delay)
		c.L.Lock()        // <8>
		queue = queue[1:] // <9>
		fmt.Println("Removed from queue")
		c.L.Unlock() // <10>
		c.Signal()   // <11>
	}

	for i := 0; i < 10; i++ {
		c.L.Lock() // <3>
		fmt.Printf("queue length %d\n", len(queue))
		for len(queue) == 5 { // <4> len(queue) == 2を満たしたらWaitに入る
			fmt.Println("do wait")
			c.Wait() // <5>
			fmt.Println("finish wait")
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) // <6>
		c.L.Unlock()                        // <7>
	}
}
