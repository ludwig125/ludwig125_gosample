package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X int
	Y int
}

func main() {
	p := &Point{X: 10, Y: 5}
	rv := reflect.ValueOf(p).Elem()
	fmt.Printf("rv.Type = %v\n", rv.Type())          // 名前空間付きの型名
	fmt.Printf("rv.Kind= %v\n", rv.Kind())           // 格納リソース種別
	fmt.Printf("rv.Interface= %v\n", rv.Interface()) // interface{}としての実際の値

	fmt.Println(rv.CanSet())
	xv := rv.Field(0)                // rv内のX要素を取り出し
	fmt.Printf("xv: %d\n", xv.Int()) // intに変換
	//xv.SetInt(100)
	//fmt.Printf("after SetInt xv: %d\n", xv.Int()) // intに変換
}
