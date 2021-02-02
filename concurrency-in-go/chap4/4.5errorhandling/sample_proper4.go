package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	type Result struct { // <1>
		Error    error
		Response *http.Response
	}
	checkStatus := func(urls ...string) <-chan Result { // <2>
		results := make(chan Result)
		for _, url := range urls {
			go func(url string) {
				defer close(results)

				fmt.Printf("http.Get %s\n", url)
				var result Result
				resp, err := http.Get(url)
				fmt.Println("sleep 2 sec")
				time.Sleep(2 * time.Second)
				result = Result{Error: err, Response: resp}
				results <- result
				fmt.Printf("result %v\n", result)
			}(url)
		}
		return results
	}

	//urls := []string{"https://www.google.com", "https://badhost", "https://www.yahoo.co.jp/"}
	urls := []string{"https://www.google.com", "https://www.yahoo.co.jp/"}
	for result := range checkStatus(urls...) {
		//		if result.Error != nil { // <5>
		//			fmt.Printf("error: %v\n", result.Error)
		//			continue
		//		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
