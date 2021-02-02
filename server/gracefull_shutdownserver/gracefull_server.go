package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("heavy process starts")
	time.Sleep(5 * time.Second)
	log.Println("done")
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("hello\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	srv := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	go func() {
		fmt.Println("starting http server on :8000")
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			log.Fatalln("Server closed with error:", err)
		}
	}()
	// http.HandleFunc("/hello", hello)
	// go func() {
	// 	if err := http.ListenAndServe(":8000", nil); err != http.ErrServerClosed {
	// 		// Error starting or closing listener:
	// 		log.Fatalln("Server closed with error:", err)
	// 	}
	// }()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		log.Println("Failed to gracefully shutdown:", err)
	}
	log.Println("Server shutdown")
}
