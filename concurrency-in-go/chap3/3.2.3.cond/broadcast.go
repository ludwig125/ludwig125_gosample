//https://github.com/kat-co/concurrency-in-go-src/blob/4e55fd7f3f5b9c5efc45a841702393a1485ba206/gos-concurrency-building-blocks/the-sync-package/cond/fig-cond-broadcast.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	type Button struct { // <1>
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) { // <2>
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup // <3>
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() { // <4>
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() { // <5>
		fmt.Println("Displaying annoying dialogue box!")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() { // <6>
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast() // <7>

	clickRegistered.Wait()
}
