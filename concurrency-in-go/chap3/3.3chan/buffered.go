package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 4)
	go func() {
		defer close(ch)
		defer fmt.Println("Producer Done")
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Printf("Sending: %d\n", i)
		}
	}()

	for i := range ch {
		fmt.Printf("Received %v\n", i)
		time.Sleep(1 * time.Second)
	}
}
