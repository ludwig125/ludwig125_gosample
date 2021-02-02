package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	if err := doTask(); err != nil {
		log.Printf("error occured. %v", err)
	}
	log.Println("finished")
}

const concurrency = 2 // 最大同時並列実行数

var errFlag bool = true

func doTask() error {
	numbers := []int{1, 2, 3, 4, 5, 6}

	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)
	errChan := make(chan error, len(numbers))
	for _, num := range numbers {
		sem <- struct{}{} // チャネルに送信
		log.Printf("num: %d", num)

		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			defer func() { <-sem }()
			log.Printf("goroutine n: %d", n)
			if err := fnA(n); err != nil {
				errChan <- fmt.Errorf("failed to A, %v", err)
				log.Printf("--> fnA len(errChan) %d", len(errChan))

				time.Sleep(1 * time.Second) // 処理をわかりやすくするため
				return
			}
			time.Sleep(1 * time.Second) // 処理をわかりやすくするため
		}(num)
	}

	go func() {
		defer close(sem)
		defer close(errChan)
		wg.Wait()
	}()

	for err := range errChan {
		return err
	}
	return nil
}

func fnA(n int) error {
	log.Println("do fnA.")
	if errFlag {
		log.Printf("--> failed to do fnA. num: %d", n)
		return fmt.Errorf("error A. num: %d", n)
	}
	log.Printf("--> succeeded to do fnA. num: %d", n)
	return nil
}
