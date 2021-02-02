package main

import "fmt"

func main() {

	st := make(chan string)
	go func() {
		st <- "abc"
	}()
	fmt.Println(<-st)
}
