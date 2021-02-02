package main

import (
	"fmt"
)

func main() {

	// 明示的な定義
	var x, y int
	x, y = 1, 2
	fmt.Println(x, y)

	var (
		n = 1
		s = "String"
		b = true
	)
	fmt.Println(n, s, b)

	// 暗黙的な定義
	x2, y2 := 10, 20
	fmt.Println(x2, y2)

	// 型指定
	n2 := 1           // int
	b2 := byte(n2)    //byte
	i64 := int64(n2)  //int64
	u32 := uint32(n2) //uint32
	fmt.Println(n2, b2, i64, u32)

}
