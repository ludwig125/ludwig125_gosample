package main

// ref: https://qiita.com/Azunyan1111/items/a1b6c58dc868814efb51

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//url := "https://www.google.com/finance/getprices?p=1d&i=3600x=TYO&q=7203" // アクセスするURLだよ！
	//url := "https://www.nikkei.com/nkd/company/history/dprice/?scode=8316&ba=1" // アクセスするURLだよ！
	//url := "https://www.nikkei.com/smartchart/?code=8316&timeframe=1d&interval=30Minute" // アクセスするURLだよ！
	//url := "https://quote.jpx.co.jp/jpx/template/qsearch.exe?F=tmp%2Fstock_list&KEY1=&KEY5=8316&KEY3=&kind=TTCODE&sort=%2B&MAXDISP=25&submit=%E6%A4%9C%E7%B4%A2%E9%96%8B%E5%A7%8B&KEY2=&KEY6=&REFINDEX=%2BTTCODE"
	url := "https://quote.jpx.co.jp/jpx/template/quote.cgi?F=tmp/stock_detail&MKTN=T&QCODE=8316"

	resp, err := http.Get(url) // GETリクエストでアクセスするよ！
	if err != nil {            // err ってのはエラーの時にエラーの内容が入ってくるよ！
		panic(err) // panicは処理を中断してエラーの中身を出力するよ！
	}
	defer resp.Body.Close() // 関数が終了するとなんかクローズするよ！（おまじない的な）

	byteArray, err := ioutil.ReadAll(resp.Body) // 帰ってきたレスポンスの中身を取り出すよ！
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteArray)) // 取り出したリクエスト結果をバイナリ配列からstring型に変換して出力するよ！
}
