package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello\n")
	}

	h := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hoge\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hello2", helloHandler)
	http.HandleFunc("/h", h)
	//log.Fatal(http.ListenAndServe("localhost:8080", nil))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
