package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"

	"github.com/sclevine/agouti"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	//log.Println("start2")

	// ユーザデータ読み取り
	flag.Parse()
	if flag.Arg(0) == "" {
		log.Fatal("Failed to get user")
	}
	user := flag.Arg(0)

	fmt.Print("Password: ")
	pass, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal("Failed to read password", err)
	} else {
		fmt.Printf("\nYour password is %v\n", string(pass))
	}

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless", // ブラウザを立ち上げないheadlessモードの指定
			//"--window-size=1280,800", // ウィンドウサイズの指定
			"--disable-gpu", // 暫定的に必要なフラグ
			"--no-sandbox",
		}),
		agouti.Debug,
	)
	if err := driver.Start(); err != nil {
		log.Printf("Failed to start driver: %v", err)
	}
	defer driver.Stop()
	log.Println("succeeded to start WebDriver")

	// WebDriverの新規セッションを作成
	page, err := driver.NewPage()
	if err != nil {
		log.Printf("Failed to open page: %v", err)
	}
	log.Println("succeeded to start new WebDriver session")

	loginURL := "https://www.k-zone.co.jp/td/users/login"
	if err := page.Navigate(loginURL); err != nil {
		log.Printf("failed to navigate: %v", err)
	}

	count(page, "sample-1")

	// IDの要素を取得し、値を設定
	identity := page.FindByID("login_id")
	if err := identity.Fill(user); err != nil {
		log.Fatalf("failed to Fill login_id: %v", err)
	}
	log.Println("succeeded to fill login")

	count(page, "sample-2")
	// passwordの要素を取得し、値を設定
	password := page.FindByName("password")
	if err := password.Fill(string(pass)); err != nil {
		log.Fatalf("failed to Fill login_id: %v", err)
	}
	log.Println("succeeded to fill pass")

	count(page, "sample-3")

	if err := page.FindByID("login_button").Submit(); err != nil {
		//if err := page.FindByButton("login_button").Submit(); err != nil {
		log.Fatalf("failed to confirm password: %v", err)
	}

	log.Println("login successfully")
}

func count(page *agouti.Page, str string) {
	log.Println(str, "find id")
	s := page.FindByID("login_button")
	//log.Printf("selection --'%#v'--, --'%v'--\n\n", sele, sele)
	//log.Printf("%T\n", s)
	cnt, err := s.Count()
	if err != nil {
		log.Fatalf("failed to select elements from %s: %v", s, err)
	}
	log.Println("len ele", cnt)
}
