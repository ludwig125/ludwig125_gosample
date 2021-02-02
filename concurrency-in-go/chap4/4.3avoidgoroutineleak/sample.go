package main

import (
	"fmt"
	"time"
)

func main() {
	dowork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("dowork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := dowork(done, nil)

	go func() {
		defer close(done)
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling dowork goroutine")
	}()
	<-terminated

}
