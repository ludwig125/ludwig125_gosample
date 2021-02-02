package main

import "fmt"

func main() {
	fmt.Println("A")
	goto L
	fmt.Println("B") // 処理されない
L: /* ラベル */
	fmt.Println("C")
}
