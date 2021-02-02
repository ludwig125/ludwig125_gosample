package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	doTask()
	log.Println("finished")
}

const concurrency = 2 // 最大同時並列実行数

func doTask() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)
	for _, num := range numbers {
		sem <- struct{}{} // チャネルに送信
		//fmt.Printf("len(sem): %d\n", len(sem))

		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			defer func() { <-sem }()
			fnA(n)

			// 処理をわかりやすくするため
			time.Sleep(1 * time.Second)
		}(num)
	}

	go func() {
		defer close(sem)
		wg.Wait()
	}()
}

func fnA(n int) {
	log.Printf("do fnA. num: %d \n", n)
}
