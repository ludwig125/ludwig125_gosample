//	Kevin Chen (2017)
//	Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

//	Golang generator pattern: functions that return channels

// https://mattn.kaoriya.net/software/lang/go/20180531104907.htm
package main

import (
	"fmt"
	"sync"
	"time"
)

func generator(s string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s %d", s, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func fanIn(ch1, ch2 <-chan string) <-chan string {
	newCh := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			newCh <- <-ch1
		}
	}()
	go func() {
		defer wg.Done()
		for {
			newCh <- <-ch2
		}
	}()
	wg.Wait()
	close(newCh)
	return newCh
}

func main() {
	f := fanIn(generator("Hello"), generator("Bye"))
	for i := range f {
		fmt.Println(i)
	}

}
