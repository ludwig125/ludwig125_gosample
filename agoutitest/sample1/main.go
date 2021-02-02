package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	// Chromeを利用することを宣言
	//agoutiDriver := agouti.ChromeDriver()
	//agoutiDriver := agouti.ChromeDriver(agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox"}))
	agoutiDriver := agouti.ChromeDriver()
	agoutiDriver.Start()
	defer agoutiDriver.Stop()
	page, _ := agoutiDriver.NewPage()

	// 自動操作
	//page.Navigate("https://talks.golang.org/2013/bestpractices.slide#1")
	page.Navigate("https://ebookjapan.yahoo.co.jp/?utm_source=yahoo&utm_medium=content&utm_campaign=pcsvc")

	// keyboard操作
	// https://github.com/sclevine/agouti/issues/61
	// https://github.com/SeleniumHQ/selenium/wiki/JsonWireProtocol#sessionsessionidelementidvalue
	time.Sleep(2 * time.Second)
	log.Println("space")
	fmt.Println("test")

	page.Screenshot("./chrome_screenshot.jpg")

	// todo
	err := page.All("html").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	for i := 0; i < 10; i++ {

		err := page.All("html").SendKeys("\uE015")
		if err != nil {
			log.Println(err)
		}
	}

	//page.MoveMouseBy(10, 100)
	//page.Find("").Click()
	//page.FindByName("").Click()
	//page.First("").Click()
	//page.Find("html").SendKeys("\uE00D")
	//	time.Sleep(2 * time.Second)
	//	page.Find("body").SendKeys("\uE00F") // PAGE_DOWN
	//	time.Sleep(2 * time.Second)
	//	//page.Find("html").SendKeys("\uE00F")
	//	time.Sleep(2 * time.Second)
	//	page.Find("body").SendKeys("\uE015") // DOWN
	//	time.Sleep(2 * time.Second)
	//	//page.Find("html").SendKeys("\uE015")

	//page.Size(1080, 800)
	//time.Sleep(2 * time.Second)
	//for i := 1000; i < 2050; i += 100 {
	//	page.MoveMouseBy(i, 100) // SPACE
	//	//time.Sleep(4 * time.Second)
	//	//page.FindByID("contents").Click() // SPACE
	//	page.First("contents").Click() // SPACE
	//	log.Println("click")
	//	time.Sleep(3 * time.Second)
	//	log.Println(i)
	//}
	time.Sleep(6 * time.Second)
}

func login(page *agouti.Page, user string, pass string, loginURL string) error {
	// ログインページに遷移
	if err := page.Navigate(loginURL); err != nil {
		return fmt.Errorf("failed to navigate: %v", err)
	}

	// html, err := page.HTML()
	// if err != nil {
	//  return fmt.Errorf("failed to get HTML: %v", err)
	// }
	// log.Println("---------------------------------------------------")
	// log.Println("HTML:", html)
	// log.Println("---------------------------------------------------")

	// HTML: view-source:https://www.k-zone.co.jp/td/users/login

	// IDの要素を取得し、値を設定
	identity := page.FindByID("login_id")
	if err := identity.Fill(user); err != nil {
		return fmt.Errorf("failed to Fill login_id: %v", err)
	}

	// passwordの要素を取得し、値を設定
	password := page.FindByName("password")
	if err := password.Fill(pass); err != nil {
		return fmt.Errorf("failed to Fill login_id: %v", err)
	}

	time.Sleep(1 * time.Second)

	//count(page, "gke_tradederby-1")

	if err := page.FindByID("login_button").Submit(); err != nil {
		//return fmt.Errorf("failed to confirm password: %v", err)
		log.Println("failed to confirm password")
	}
	return nil
}
