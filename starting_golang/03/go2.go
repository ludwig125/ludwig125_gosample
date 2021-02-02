package main

import (
	"fmt"
	"runtime"
)

func main() {
	go fmt.Println("Yeah") // go 文を追加してみる
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %s\n", runtime.Version())
}
