package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	type Result struct {
		Error    error
		Response *http.Response
	}
	checkStatus := func(urls []string) <-chan Result {
		resultChan := make(chan Result, 10)

		for _, url := range urls {
			go func(url string) {
				defer close(resultChan)
				resp, err := http.Get(url)
				fmt.Println("sleep 2 sec")
				time.Sleep(2 * time.Second)
				resultChan <- Result{Error: err, Response: resp}
			}(url)
		}
		return resultChan
	}
	urls := []string{"https://www.google.com", "https://badhost", "https://www.yahoo.co.jp/"}

	for result := range checkStatus(urls) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
