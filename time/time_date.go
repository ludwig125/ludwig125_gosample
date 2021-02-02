package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func getDate(s string) string {
	re := regexp.MustCompile(`\d+:\d+`).Copy()
	t := strings.Split(re.FindString(s), ":") // ["06", "00"]
	data_hour, _ := strconv.Atoi(t[0])
	data_hour = data_hour + 9 // GMT -> JST
	data_min := t[1]

	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst)

	d := now.Weekday()
	h := now.Hour()

	var ymd string
	switch d {
	case 1: // Monday
		if h < data_hour {
			// 月曜に取得したデータが現在時刻より後であればそれは３日前の金曜のデータ
			ymd = now.AddDate(0, 0, -3).Format("2006/01/02")
		} else {
			ymd = now.Format("2006/01/02")
		}
	case 2, 3, 4, 5: // Tuesday,..Friday
		if h < data_hour {
			// 火~金曜に取得したデータが現在時刻より後であればそれは前日のデータ
			ymd = now.AddDate(0, 0, -1).Format("2006/01/02")
		} else {
			ymd = now.Format("2006/01/02")
		}
	case 6: // Saturday
		// 土曜に取得したデータは前日の金曜のデータ
		ymd = now.AddDate(0, 0, -1).Format("2006/01/02")
	case 0: // Sunday
		// 日曜に取得したデータは２日前の金曜のデータ
		ymd = now.AddDate(0, 0, -2).Format("2006/01/02")
	}
	return fmt.Sprintf("%s %2d:%s", ymd, data_hour, data_min)
}

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

	fmt.Println(getDate("06:00 "))
}
