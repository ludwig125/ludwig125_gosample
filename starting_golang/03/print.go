package main

import (
	"fmt"
)

func main() {
	n := 1
	fmt.Println(n)

	fmt.Printf("数値=%d\n", 5)

	// %vはいろいろな型に対応
	fmt.Printf("数値=%v 文字列=%v 配列=%v\n", 5, "Golang Go言語", [...]int{1, 2, 3})
}
