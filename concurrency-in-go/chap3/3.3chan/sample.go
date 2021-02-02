package main

import "fmt"

func main() {
	c := make(chan string, 3)

	go func() {
		defer close(c)
		for i := 0; i < 3; i++ {
			c <- fmt.Sprintf("%d desu", i)
		}
	}()
	for x := range c {
		fmt.Println(x)
	}
}
