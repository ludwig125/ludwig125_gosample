package main

import (
	"fmt"
	"runtime"
	"time"
)

// https://blog.gopheracademy.com/advent-2018/avoid-gc-overhead-large-heaps/

func main() {
	a := make([]*int, 1e9)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s\n", time.Since(start))
	}

	fmt.Println("len a", len(a))
	runtime.KeepAlive(a)
}

// func main() {
// 	a := make([]int, 1e9)

// 	for i := 0; i < 10; i++ {
// 		start := time.Now()
// 		runtime.GC()
// 		fmt.Printf("GC took %s\n", time.Since(start))
// 	}

// 	runtime.KeepAlive(a)
// }
