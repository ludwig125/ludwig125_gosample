package main

import (
	"fmt"
)

func main() {
	var x = 123

	c := make(chan struct{})

	go func() {
		x = 789 // write x
		c <- struct{}{}
	}()

	<-c
	fmt.Println(x) // read x
}
