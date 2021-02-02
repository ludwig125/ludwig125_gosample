//	Kevin Chen (2017)
//	Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

//	Golang generator pattern: functions that return channels

// https://mattn.kaoriya.net/software/lang/go/20180531104907.htm
package main

import (
	"fmt"
	"time"
)

// goroutine is launched inside the called function (more idiomatic)
// multiple instances of the generator may be called

func generator(s string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i <= 5; i++ {
			ch <- fmt.Sprintf("%s %d", s, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func main() {
	g := generator("Hello")
	for i := 0; i <= 5; i++ {
		fmt.Println(<-g)
	}

	for v := range generator("Hello") {
		fmt.Println(v)
	}
}
