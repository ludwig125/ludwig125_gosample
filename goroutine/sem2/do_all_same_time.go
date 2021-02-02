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

func doTask() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	var wg sync.WaitGroup
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fnA(n)

			// 処理をわかりやすくするため
			time.Sleep(1 * time.Second)
		}(num)
	}
	wg.Wait()
}

func fnA(n int) {
	log.Printf("do fnA. num: %d \n", n)
}
