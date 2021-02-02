// 参考：http://gihyo.jp/dev/feature/01/go_4beginners/0005?page=3
package main

import (
	"fmt"
	"log"
	"net/http"
)

func getStatus(urls []string) <-chan string {
	statusChan := make(chan string)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			statusChan <- res.Status
		}(url)
	}
	return statusChan
}

func main() {
	urls := []string{"https://www.google.com", "https://badhost", "https://www.yahoo.co.jp/"}
	//urls := []string{"https://www.google.com", "https://www.yahoo.co.jp/"}
	statusChan := getStatus(urls)

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}
