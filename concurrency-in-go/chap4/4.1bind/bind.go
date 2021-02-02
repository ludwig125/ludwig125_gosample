package main

import (
	"fmt"
)

func main() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5) // このchanには他のGoroutineが書き込めない
		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) { // intのchanの読み込み専用のコピーを受け取る。読み込み権限のみが必要
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
	}

	results := chanOwner() // chanの読み込み権限を受け取って消費者に渡している。読み込み以外は何もしない
	consumer(results)

}
