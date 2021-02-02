package main

import (
	"fmt"
	"net/http"
)

func jsonResponse(rw http.ResponseWriter, req *http.Request) {
	// response := []byte(`
	// [
	//   	{"name":"a"},
	// 	{"name":"b"},
	//  	{"name":"c"},
	//  	{"name":"d"}
	// ]`)
	response := `
	[
	  	{"name":"aaa"},
  		{"name":"bbb"},
 	 	{"name":"ccc"},
 	 	{"name":"ddd"}
	]
	`

	defer func() {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, string(response))
	}()
}

func main() {
	http.HandleFunc("/json", jsonResponse)
	http.ListenAndServe(":8080", nil)
}
