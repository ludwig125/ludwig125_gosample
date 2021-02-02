// 参考：http://gihyo.jp/dev/feature/01/go_4beginners/0005?page=3
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	checkStatus := func(urls []string) <-chan string {
		resultChan := make(chan string)
		for _, url := range urls {
			go func(url string) {
				resp, err := http.Get(url)
				fmt.Println("sleep 2 sec")
				time.Sleep(2 * time.Second)
				if err != nil {
					fmt.Printf("http.Get error: %v\n", err)
				}
				resultChan <- resp.Status
			}(url)
		}
		return resultChan
	}
	//urls := []string{"https://www.google.com", "https://badhost", "https://www.yahoo.co.jp/"}
	urls := []string{"https://www.google.com", "https://www.yahoo.co.jp"}
	resultChan := checkStatus(urls)

	for i := 0; i < len(urls); i++ {
		fmt.Printf("Response: %s\n", <-resultChan)
	}
}
