package main

import "fmt"

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int, 10)

	// チャネルに整数5を送信
	ch <- 5
	// チャネルから整数値を受信
	i := <-ch
	fmt.Println(i)

	// 新しくチャネルを作って関数とやりとりする
	ch1 := make(chan int)

	go receiver(ch1)

	j := 0
	for j < 1000 {
		ch1 <- j
		j++
	}

	//num, ok := <-ch1
	//fmt.Printf("%d %v", num, ok)
	close(ch1)
	num, ok := <-ch1
	fmt.Printf("%d %v", num, ok)
}
