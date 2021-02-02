package main

import (
	"fmt"
	"net/http"
)

// StatusHandler is struct.
type StatusHandler struct {
	ReqCnt int
}

func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.ReqCnt++
	content := fmt.Sprintf(`{"status":"ok", "count": %d}`, h.ReqCnt)
	w.Write([]byte(content))
}

// StaticHandler is struct.
type StaticHandler struct{}

func (h *StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"ok"}`))
}

func main() {
	status := StatusHandler{}
	static := StaticHandler{}
	mux := http.NewServeMux()
	mux.Handle("/status", &status)
	mux.Handle("/static", &static)

	http.ListenAndServe("localhost:9000", mux)
}
