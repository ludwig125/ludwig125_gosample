package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type result struct {
	Error error
	Res   int
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	queue := make(chan string)
	res := make(chan result, 3)

	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(queue chan string) {
			for {
				select {
				case <-ctx.Done():
					defer wg.Done()
					fmt.Println("worker exit")
					return
				case url := <-queue:
					st, err := getBody(url)
					//fmt.Println(st, err)
					res <- result{err, st}
				}
			}
		}(queue)
	}
	queue <- "https://www.yahoo.co.jp/"
	queue <- "https://www.google.co.jp/"
	queue <- "https://www.google2.co.jp/"

	cancel()
	wg.Wait()
	close(res)
	for r := range res {
		log.Printf("res: %d err: %v\n", r.Res, r.Error)
	}

	fmt.Println("------------")
	doTask()
}

func doTask() {
	ctx, cancel := context.WithCancel(context.Background())
	queue := make(chan string)
	res := make(chan int, 3)
	errCh := make(chan error, 3)

	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(queue chan string) {
			for {
				select {
				case <-ctx.Done():
					defer wg.Done()
					fmt.Println("worker exit")
					return
				case url := <-queue:
					st, err := getBody(url)
					//fmt.Println(st, err)
					res <- st
					errCh <- err
				}
			}
		}(queue)
	}
	queue <- "https://www.yahoo.co.jp/"
	queue <- "https://www.google.co.jp/"
	queue <- "https://www.google2.co.jp/"

	cancel()
	wg.Wait()
	close(res)
	close(errCh)
	for r := range res {
		log.Printf("res: %d\n", r)
	}
	for err := range errCh {
		log.Printf("err: %v\n", err)
	}
}

func getBody(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to http.Get, %v", err)
	}

	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	return resp.StatusCode, nil
}
