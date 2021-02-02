package main

import (
	"fmt"
	"time"
)

func main() {
	var x = 123

	go func() {
		x = 789 // write x
	}()

	time.Sleep(time.Second)
	fmt.Println(x) // read x
}
