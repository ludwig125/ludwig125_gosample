package main

import (
	"fmt"
	"log"

	"github.com/sclevine/agouti"
)

func main() {
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
			//"--window-size=300,400",
			//"--blink-settings=imagesEnabled=false", // don't load images
			"--disable-gpu", // ref: https://developers.google.com/web/updates/2017/04/headless-chrome#cli
			"--no-sandbox",  // ref: https://github.com/theintern/intern/issues/878
			//"--disable-dev-shm-usage", // ref: https://qiita.com/yoshi10321/items/8b7e6ed2c2c15c3344c6
			"--whitelisted-ips", //https://chromedriver.chromium.org/security-considerations
		}),
		agouti.Debug,
	)
	if err := driver.Start(); err != nil {
		log.Fatal(err)
	}
	defer driver.Stop()
	page, err := driver.NewPage()
	if err != nil {
		log.Fatal(err)
	}
	page.Navigate("https://golang.org/")
	getSource, err := page.HTML() // htmlソースを取得
	if err != nil {
		log.Fatal("get HTML", err)
	}
	fmt.Println(getSource)
	page.Screenshot("Screen.png") // スクリーンショット
}

// package main

// import (
// 	"log"
// 	"os"

// 	"github.com/sclevine/agouti"
// )

// func main() {
// 	// ブラウザはChromeを指定して起動
// 	//driver := agouti.ChromeDriver(agouti.Browser("chrome"))
// 	driver := agouti.ChromeDriver(
// 		agouti.ChromeOptions("args", []string{
// 			//		"--headless",             // headlessモードの指定
// 			"--window-size=1280,800", // ウィンドウサイズの指定
// 		}),
// 		agouti.Debug,
// 	)
// 	if err := driver.Start(); err != nil {
// 		log.Fatalf("Failed to start driver: %v", err)
// 		os.Exit(0)
// 	}
// 	defer driver.Stop()
// 	log.Println("started chromedriver successfully")

// 	page, err := driver.NewPage(agouti.Browser("chrome"))
// 	if err != nil {
// 		log.Fatalf("Failed to open page:%v", err)
// 	}

// 	if err := page.Navigate("http://qiita.com/"); err != nil {
// 		log.Fatalf("Failed to navigate:%v", err)
// 	}
// 	page.Screenshot("/tmp/chrome_qiita.jpg")
// 	// page, err := driver.NewPage()
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to open page: %v", err)
// 	// 	os.Exit(0)
// 	// }
// 	// log.Println("created NewPage successfully")

// 	// time.Sleep(2 * time.Second)
// 	// page.Screenshot("./chrome_screenshot.jpg")
// }
