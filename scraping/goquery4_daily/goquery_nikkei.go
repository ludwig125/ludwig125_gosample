// get stock price and time for code 8316
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape(code string) {
	// Request the HTML page.
	url := "https://gae-webui.appspot.com?code=" + code
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

	var dp [][]string
	// Find time and yen
	doc.Find(".m-tableType01_table table tbody tr").Each(func(i int, s *goquery.Selection) {
		date := s.Find(".a-taC").Text()
		re := regexp.MustCompile(`[0-9]+/[0-9]+`).Copy()
		date = re.FindString(date)

		var p []string
		p = append(p, date)
		s.Find(".a-taR").Each(func(i2 int, s2 *goquery.Selection) {
			//fmt.Println("test", date, s2.Text())
			p = append(p, strings.Replace(s2.Text(), ",", "", -1))
		})
		//fmt.Println(p)
		dp = append(dp, p)
		//date_price[i] = p

	})
	fmt.Println(dp)
	fmt.Println(dp[0][0], dp[0][1], dp[0][2], dp[0][3])

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
