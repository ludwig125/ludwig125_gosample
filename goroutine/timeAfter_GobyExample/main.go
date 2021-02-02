package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	select {
	case res := <-task():
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
	log.Printf("Total time: %fs", time.Since(start).Seconds())
}

// func main() {
// 	start := time.Now()
// 	c1 := make(chan string, 1)
// 	go func() {
// 		<-task()
// 		c1 <- "result 1"
// 	}()
// 	select {
// 	case res := <-c1:
// 		fmt.Println(res)
// 	case <-time.After(1 * time.Second):
// 		fmt.Println("timeout 1")
// 	}
// 	log.Printf("Total time: %fs", time.Since(start).Seconds())
// }

func task() <-chan string {
	res := make(chan string)
	go func() {
		defer close(res)
		for i := 0; i < 1000000; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Nanosecond)
		}
		res <- "done!"
	}()
	return res
}
