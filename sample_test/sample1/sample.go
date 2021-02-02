package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func remoteHello(domain string) string {
	// /greetingにクエリパラメータgreet=Helloを渡してGet問い合わせする
	res, err := http.Get(domain + "/greeting?greet=Hello")

	// エラー処理
	if err != nil {
		fmt.Println("Error")
		return "error"
	}
	defer res.Body.Close()

	// レスポンスを戻り値にする
	res_str, _ := ioutil.ReadAll(res.Body)
	return string(res_str)
}
