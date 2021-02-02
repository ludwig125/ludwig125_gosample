package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		fmt.Println("start..")
		time.Sleep(3 * time.Second)
		close(c)
	}()

	select {
	case <-c:
		fmt.Printf("get after %v\n", time.Since(start))
	}
}
