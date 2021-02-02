package main

import (
	"fmt"
	"net/http"
)

func jsonResponse(rw http.ResponseWriter, req *http.Request) {
	response := `[
{"id":"1000000001.aaa","fields":{"account":1000000001,"item":"aaa","local":1001}},
{"id":"1000000002.bbb","fields":{"account":1000000002,"item":"bbb","local":1002}},
{"id":"1000000003.ccc","fields":{"account":1000000003,"item":"ccc","local":1003}}]
`
	//response := `[]`
	defer func() {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, string(response))
	}()
}

func main() {
	http.HandleFunc("/json", jsonResponse)
	http.ListenAndServe(":8080", nil)
}
