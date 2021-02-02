// 伝統的な `printf` によく似た、良く出来た文字列フォーマット機能が Go にはある。
// ここでは、文字列をフォーマットする例をいくつか紹介する。

package main

import (
	"fmt"
)

func main() {

	// 文字列を普通に表示するには `%s` を使う。
	fmt.Printf("%s\n", "\"string\"")

	// Go のソースコードのようにダブルクオートを文字列に入れるには、`%q` を使う。
	fmt.Printf("%q\n", "\"string\"")

}
