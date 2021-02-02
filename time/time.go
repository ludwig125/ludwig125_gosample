package main

import (
	"fmt"
	"time"
)

func main() {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst)
	fmt.Println(now) // 2018-08-08 23:49:03.599423499 +0900 JST
	// goの日付フォーマットは独特で2006-01-02のように指定する
	// https://golang.org/pkg/time/#Time.In
	// https://blog.toshimaru.net/go-time-format/
	// http://text.baldanders.info/golang/time-functions/
	// https://ashitani.jp/golangtips/tips_time.html
	fmt.Println(now.Format("2006/01/02")) // 2018/08/08

}
