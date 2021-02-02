// get stock price and time for code 8316
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://www.nikkei.com/smartchart/?code=8316&timeframe=1d&interval=30Minute")
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

	// Find the review items
	doc.Find(".stockInfoinner").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//no := s.Find(".no span").First().Text()
		//name := s.Find(".name").Text()
		time := s.Find(".ttl1").Text()
		yen := s.Find(".item1").Text()
		//fmt.Println(no, name, time, yen)
		fmt.Println(time, yen)
	})
}

func main() {
	ExampleScrape()
}
