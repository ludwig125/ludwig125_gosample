// get stock price and time for code 8316
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape(code string) {
	// Request the HTML page.
	url := "https://www.nikkei.com/smartchart/?code=" + code
	//res, err := http.Get("https://www.nikkei.com/smartchart/?code=8316&timeframe=1d&interval=30Minute")
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find time and yen
	doc.Find(".stockInfoinner").Each(func(i int, s *goquery.Selection) {
		//no := s.Find(".no span").First().Text()
		//name := s.Find(".name").Text()
		time := s.Find(".ttl1").Text()
		yen := s.Find(".item1").Text()
		//fmt.Println(no, name, time, yen)
		fmt.Println(code, time, yen)
	})
}

func main() {
	file, err := os.Open(`./code`)
	if err != nil {
		// Openエラー処理
		fmt.Fprintf(os.Stderr, "File %s could not read")
		os.Exit(1)
	}
	defer file.Close()

	// Scannerで読み込む
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ExampleScrape(scanner.Text())
	}

}
