package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan interface{})
	go func() {
		defer close(done)
		fmt.Println("Goroutine start")
		time.Sleep(5 * time.Second)
		fmt.Println("Goroutine end. send done")
	}()

	i := 0
	for {
		select {
		case <-done:
			fmt.Println("Receive done")
			return
		default:
		}

		i++
		fmt.Printf("Wait %d\n", i)
		time.Sleep(1 * time.Second)
	}

}
