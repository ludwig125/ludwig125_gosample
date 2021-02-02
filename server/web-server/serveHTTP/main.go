package main

import "net/http"

// CustomHandler is struct.
type CustomHandler struct{}

func (hand *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("custom handler!"))
}

func main() {
	hand := CustomHandler{}
	mux := http.NewServeMux()
	mux.Handle("/", &hand)

	http.ListenAndServe("localhost:3000", mux)
}
