package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a []int // nil
	var b bool  // false

	// a new goroutine
	go func() {
		a = make([]int, 3)
		b = true // write b
	}()

	for !b { // read b
		time.Sleep(time.Second)
		runtime.Gosched()
	}
	a[0], a[1], a[2] = 0, 1, 2 // might panic
	for _, v := range a {
		fmt.Println(v)
	}
}
