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
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result { // <2>
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				fmt.Printf("http.Get %s\n", url)
				var result Result
				resp, err := http.Get(url)
				fmt.Println("sleep 2 sec")
				time.Sleep(2 * time.Second)
				result = Result{Error: err, Response: resp} // <3>
				select {
				case <-done:
					return
				case results <- result: // <4>
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost", "https://www.yahoo.co.jp/"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil { // <5>
			fmt.Printf("error: %v\n", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
