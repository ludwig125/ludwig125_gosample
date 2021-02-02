// http://cuto.unirita.co.jp/gostudy/post/standard-library-file-io/
package main

import (
	"fmt"
	"os"
)

const BUFSIZE = 1024 // 読み込みバッファのサイズ

func main() {
	file, err := os.Open(`./code`)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	buf := make([]byte, BUFSIZE)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			// Readエラー処理
			break
		}

		fmt.Print(string(buf[:n]))
	}
}
