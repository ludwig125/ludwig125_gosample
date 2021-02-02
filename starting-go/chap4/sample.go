package main

import "fmt"

func main() {
	//ch := make(chan int, 1)
	ch := make(chan int)
	go func() {
		ch <- 5
	}()

	fmt.Println(<-ch)
}
