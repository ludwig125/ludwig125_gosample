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
	// log.Println("heavy process starts")
	// time.Sleep(5 * time.Second)
	// log.Println("done")
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

	// go func() {
	// 	fmt.Println("starting http server on :8000")
	// 	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
	// 		// Error starting or closing listener:
	// 		log.Fatalln("Server closed with error:", err)
	// 	}
	// }()
	serverErrCh := server(srv)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	closeServerErrCh := closeServer(ctx, srv)
	<-serverErrCh
	log.Println("serverErrCh")
	<-closeServerErrCh
	log.Println("closeServerErrCh")
	// select {
	// case <-serverErrCh:
	// 	log.Println("serverErrCh")
	// case <-closeServerErrCh:
	// 	log.Println("closeServerErrCh")
	// }
	// if err := srv.Shutdown(ctx); err != nil {
	// 	// Error from closing listeners, or context timeout:
	// 	log.Println("Failed to gracefully shutdown:", err)
	// }
	// log.Println("Server shutdown")
}

func server(srv *http.Server) chan error {
	errCh := make(chan error)
	go func() {
		defer close(errCh)
		fmt.Println("starting http server on :8000")
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			log.Println("Server closed with error:", err)
			errCh <- err
			return
		}
		log.Println("Server closed successfully")
	}()
	return errCh
}

func closeServer(ctx context.Context, srv *http.Server) chan error {
	errCh := make(chan error)
	go func() {
		defer close(errCh)
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Println("Failed to gracefully shutdown:", err)
			errCh <- err
			return
		}
		log.Println("Server shutdown gracefully")
	}()
	return errCh
}
