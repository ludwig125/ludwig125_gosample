// https://qiita.com/Azunyan1111/items/a1b6c58dc868814efb51
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World") // Hello, Worldってアクセスした人に返すよ！
}

func main() {
	http.HandleFunc("/", handler)     // http://localhost:8080/にアクセスしてきた人はhandlerを実行する
	http.ListenAndServe(":8080", nil) // サーバーを起動する
}
