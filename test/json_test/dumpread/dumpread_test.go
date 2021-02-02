package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

var (
	response = `[
{"id":"1000000001.aaa","fields":{"account":1000000001,"item":"aaa","local":1001}},
{"id":"1000000002.bbb","fields":{"account":1000000002,"item":"bbb","local":1002}},
{"id":"1000000003.ccc","fields":{"account":1000000003,"item":"ccc","local":1003}}]
`
)

func dummyHandler(t *testing.T) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/json" {
			t.Fatalf("invalid access path %v", r.URL.Path)
		}
		if r.URL.Query().Get("type") != "documents" {
			t.Fatalf("invalid params")
		}
		w.Header().Set("content-Type", "application/json")
		fmt.Fprintf(w, string(response))
		return
	})
}

func TestGetBody(t *testing.T) {
	// ts := httptest.NewServer(http.HandlerFunc(
	// 	func(w http.ResponseWriter, r *http.Request) {
	// 		if r.URL.Path != "/json" {
	// 			t.Fatalf("invalid access path %v", r.URL.Path)
	// 		}
	// 		if r.URL.Query().Get("type") != "documents" {
	// 			t.Fatalf("invalid params")
	// 		}
	// 		w.Header().Set("content-Type", "application/json")
	// 		fmt.Fprintf(w, string(response))
	// 		return
	// 	},
	// ))
	ts := httptest.NewServer(dummyHandler(t))
	defer ts.Close()

	t.Run("getBody", func(t *testing.T) {
		testGetBodySuccess(t, ts.URL)
		testWriteFile(t, ts.URL)
		testJSONParse(t, ts.URL)
	})
}

func testGetBodySuccess(t *testing.T, url string) {
	resp, err := getBody(url)
	if err != nil {
		t.Fatal(err)
	}
	io.Copy(ioutil.Discard, resp)
	resp.Close()
}

func testWriteFile(t *testing.T, url string) {
	if err := writeFile(url); err != nil {
		t.Fatal(err)
	}
	data, err := ioutil.ReadFile(dumpFile)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != response {
		t.Fatal("file data unmatch")
	}
}

func testJSONParse(t *testing.T, url string) {
	_ = writeFile(url)

	f, _ := os.Open(dumpFile)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var decodedID []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "[") {
			continue
		}
		//fmt.Println(line)
		got, err := jsonParse(line)
		if err != nil {
			t.Fatal(err)
		}
		// fmt.Println(decoded.ID)
		// fmt.Println(decoded.Fields)
		decodedID = append(decodedID, got.ID)
	}
	want := []string{"1000000001.aaa", "1000000002.bbb", "1000000003.ccc"}
	//fmt.Println(decodedID)
	if !reflect.DeepEqual(want, decodedID) {
		t.Fatal("document id unmatch")
	}
}

// func ExampleJsonParse() {
// 	//_ = writeFile(url)

// 	f, _ := os.Open(dumpFile)
// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if strings.HasPrefix(line, "[") {
// 			continue
// 		}
// 		got, _ := jsonParse(line)
// 		fmt.Println(got.ID)
// 	}
// 	// Unordered output: 3
// 	// 2
// 	// 1
// 	// 3
// }
