package main

import "fmt"

func main() {
	hello := func() {
		fmt.Println("hello")
	}
	go hello()
}
