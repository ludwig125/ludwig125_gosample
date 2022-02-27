package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// StatusHandler is struct.
type StatusHandler struct {
	ReqCnt int
}

func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.ReqCnt++
	content := fmt.Sprintf(`{"status":"ok", "count": %d}`, h.ReqCnt)
	log.Println("called status")

	// Atrificially delay this function to emulate a long running task
	time.Sleep(1 * time.Second)

	w.Write([]byte(content))
}

// StaticHandler is struct.
type StaticHandler struct{}

func (h *StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("called static")
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
