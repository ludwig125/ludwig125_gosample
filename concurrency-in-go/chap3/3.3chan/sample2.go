package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan interface{})
	for i := 0; i < 3; i++ {
		go func() {
			<-c
			fmt.Println(i)
		}()
	}
	close(c)
}
