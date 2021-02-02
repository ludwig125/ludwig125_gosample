package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func getTime(s string) string {
	re := regexp.MustCompile(`\d+:\d+`).Copy()
	t := strings.Split(re.FindString(s), ":") // ["06", "00"]
	hour, _ := strconv.Atoi(t[0])
	min := t[1]

	// 以下で年月日を取得
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst)
	ymd := now.Format("2006/01/02")
	return fmt.Sprintf("%s %2d:%s", ymd, hour+9, min)
}

func getYen(s string) string {
	re := regexp.MustCompile(`[0-9,.]+`).Copy()
	yen := re.FindString(s)
	//yen = strings.TrimRight(yen, ".0")
	yen = strings.Replace(yen, ".0", "", 1)
	yen = strings.Replace(yen, ",", "", -1)
	return yen
}

func main() {
	fmt.Println(getTime("現在値(06:00)："))

	fmt.Println(getYen("680.0円"))
	fmt.Println(getYen("4,415円"))
	fmt.Println(getYen("4,415.0円"))
	fmt.Println(getYen("1,124,415.0円"))
}
