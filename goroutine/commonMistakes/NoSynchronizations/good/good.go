package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a []int // nil
	var b bool  // false

	waitCh := make(chan interface{})

	// a new goroutine
	go func() {
		a = make([]int, 3)
		b = true // write b
		waitCh <- struct{}{}
	}()

	<-waitCh

	for !b { // read b
		time.Sleep(time.Second)
		runtime.Gosched()
	}
	a[0], a[1], a[2] = 0, 1, 2 // would not panic
	for _, v := range a {
		fmt.Println(v)
	}
}
