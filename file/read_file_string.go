// https://qiita.com/jpshadowapps/items/ae7274ec0d40882d76b5
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//lines := fromFile("./code")
	//fmt.Printf("lines: %v\n", lines)
	fromFile("./code")
}

//func fromFile(filePath string) []string {
func fromFile(filePath string) {
	// ファイルを開く
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filePath, err)
		os.Exit(1)
	}

	// 関数return時に閉じる
	defer f.Close()

	// Scannerで読み込む
	// lines := []string{}
	//lines := make([]string, 0, 100)  // ある程度行数が事前に見積もれるようであれば、makeで初期capacityを指定して予めメモリを確保しておくことが望ましい
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// appendで追加
		//lines = append(lines, scanner.Text())
		fmt.Println(scanner.Text())
	}
	if serr := scanner.Err(); serr != nil {
		fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", filePath, err)
	}

	return
	//return lines
}
