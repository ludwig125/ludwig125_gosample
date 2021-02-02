package main

import "fmt"

func main() {
	ch := make(chan int, 10)

	// チャネルに５を送信
	ch <- 5

	// チャネルから整数値を受信
	i := <-ch
	fmt.Println(i)
}
