package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteHello(t *testing.T) {
	// テストサーバを用意する
	// サーバ側でアクセスする側のテストを行う
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// URLのアクセスパスが誤っていないかチェック
			if r.URL.Path != "/greeting" {
				t.Fatalf("誤ったアクセスパスでアクセス!")
			}
			// クエリパラメータをチェック
			if r.URL.Query().Get("greet") != "Hello" {
				t.Fatalf("正しく挨拶してない!")
			}
			// レスポンスを設定する
			w.Header().Set("content-Type", "text")
			fmt.Fprintf(w, "world")
			return
		},
	))
	defer ts.Close()

	// クライアントのコードを呼び出す.
	// アクセスされる側(サーバ)のレスポンスのテストを行う.

	// テストサーバのルートパスはts.URLで取得できます
	res := remoteHello(ts.URL)

	if res != "world" {
		t.Errorf("世界じゃなかった.")
	}
}
