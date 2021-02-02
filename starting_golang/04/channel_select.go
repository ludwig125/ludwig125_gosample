package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	/* ch1から受信した整数を２倍してch2へ送信 */
	go func() {
		for {
			i := <-ch1
			ch2 <- (i * 2)
		}
	}()

	/* ch2から受信した整数を１減算してch3へ送信*/
	go func() {
		for {
			i := <-ch2
			ch3 <- (i - 1)
		}
	}()

	n := 1
LOOP:
	for {
		select {
		case ch1 <- n:
			n++
		case i := <-ch3:
			fmt.Println("received", i)
		default:
			if n > 100 {
				break LOOP
			}
		}
	}
}
