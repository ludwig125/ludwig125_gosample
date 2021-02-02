package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

type stockInfo struct {
	Code   string
	Name   string
	Status string
}

func main() {
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
		fmt.Printf("\nYour password is %v\n", string(password))
	}

	html, err := fetchStockDocFromWebPage(user, pass)
	if err != nil {
		log.Fatalf("Failed to fetchStockDocFromWebPage, %v", err)
	}

	stockInfos, err := fetchStockInfo(html)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stockInfos)
}
