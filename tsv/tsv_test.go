package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

type testdatabase struct {
}

// interface を満たしているか確認する
var _ DB = (*database)(nil)

func NewTestDB() DB {
	return testdatabase{}
}

func (d testdatabase) ReadPartition(table string) (string, error) {
	res := `
PART_NAME
partition=20211021
partition=20211022
partition=20211023
partition=20211024`
	return res, nil
}

func (d testdatabase) GetTSV() (string, error) {
	return "ColumnA\tColumnB\na1\tb1\na2\tb2\na3\tb3\n", nil
}

type testserver struct {
	serverURL string
}

// interface を満たしているか確認する
var _ Server = (*testserver)(nil)

func NewTestServer(serverURL string) Server {
	return testserver{serverURL: serverURL}
}

func (s testserver) Delete(params string) error {
	return delete(s.serverURL, params)
}

func (s testserver) Read(params string) (string, error) {
	return "", nil
}

func TestRun(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if r.Header.Get("Content-Type") != "application/json" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			//To allocate slice for request body
			length, err := strconv.Atoi(r.Header.Get("Content-Length"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			body := make([]byte, length)
			length, err = r.Body.Read(body)
			if err != nil && err != io.EOF {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			//parse json
			var jsonBody map[string]interface{}
			if err := json.Unmarshal(body[:length], &jsonBody); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fmt.Printf("%v\n", jsonBody)

			// params := r.FormValue("parameter1")

			// // URLのアクセスパスが誤っていないかチェック
			// if r.URL.Path != "/greeting" {
			// 	t.Fatalf("誤ったアクセスパスでアクセス!")
			// }
			// // クエリパラメータをチェック
			// if r.URL.Query().Get("greet") != "Hello" {
			// 	t.Fatalf("正しく挨拶してない!")
			// }
			// レスポンスを設定する
			w.Header().Set("content-Type", "text")
			res := fmt.Sprintf("res: %v", jsonBody)
			fmt.Fprintf(w, string(res))
		},
	))
	defer ts.Close()

	db := NewTestDB()
	srv := NewTestServer(ts.URL)
	if err := run(db, time.Date(2021, 10, 24, 1, 2, 3, 4, time.Local), srv); err != nil {
		t.Fatal(err)
	}

	// tests := map[string]struct {
	// 	data    string
	// 	wantErr bool
	// 	want    []Data
	// }{
	// 	"1": {
	// 		data: "a1\tb1\na2\tb2\na3\tb3\n",
	// 		want: []Data{
	// 			{A: "a1", B: "b1"},
	// 			{A: "a2", B: "b2"},
	// 			{A: "a3", B: "b3"},
	// 		},
	// 	},
	// }
	// for name, tt := range tests {
	// 	t.Run(name, func(t *testing.T) {
	// 		fmt.Println(tt)
	// 		// got, err := ConvertDataFromTSV(strings.NewReader(tt.data))
	// 		// if err != nil {
	// 		// 	if !tt.wantErr {
	// 		// 		t.Errorf("error: %v, wantErr: %t", err, tt.wantErr)
	// 		// 	}
	// 		// 	t.Log(err)
	// 		// 	return
	// 		// }

	// 		// if diff := cmp.Diff(got, tt.want); diff != "" {
	// 		// 	t.Errorf("got: %v,want: %v, diff: %s", got, tt.want, diff)
	// 		// }
	// 	})
	// }
}

func TestReadTSV(t *testing.T) {
	tests := map[string]struct {
		data    string
		wantErr bool
		want    []Data
	}{
		"1": {
			data: "ColumnA\tColumnB\na1\tb1\na2\tb2\na3\tb3\n",
			want: []Data{
				{A: "a1", B: "b1"},
				{A: "a2", B: "b2"},
				{A: "a3", B: "b3"},
			},
		},
		"empty": {
			data:    "",
			wantErr: false,
			want:    nil,
		},
		"abnormal": {
			data:    "ColumnA\tColumnB\na1\tb1\na2",
			wantErr: true, // error occured in read tsv: record on line 3: wrong number of fields
			want:    nil,
		},
		"abnormal2": {
			data:    "ColumnA\tColumnB\na1\na2",
			wantErr: true, // tsv data does not have exactly two fields. data: [a1]
			want:    nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ConvertDataFromTSV(strings.NewReader(tt.data))
			if err != nil {
				if !tt.wantErr {
					t.Errorf("error: %v, wantErr: %t", err, tt.wantErr)
				}
				t.Log(err)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("got: %v,want: %v, diff: %s", got, tt.want, diff)
			}
		})
	}
}
