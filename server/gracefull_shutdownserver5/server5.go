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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan interface{})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}
	serverErrCh := server(done, srv)
	closeServerErrCh := doneServer(ctx, srv, done)

	// do something

	select {
	case s := <-sig:
		log.Printf("signal %d received", s)
		close(done)
	case <-ctx.Done():
		log.Printf("context canceled: %v", ctx.Err())
		close(done)
	case <-done:
		log.Println("done received")
	}

	select {
	case err := <-serverErrCh:
		if err != nil {
			log.Printf("serverErr: %v", err)
		}
	case err := <-closeServerErrCh:
		if err != nil {
			log.Printf("closeServerErrCh: %v", err)
		}
		//// serverのcloseも待つ
		//<-serverErrCh
	}
	//time.Sleep(3 * time.Second)
	log.Print("Server Stopped")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func server(done chan interface{}, srv *http.Server) chan error {
	errCh := make(chan error)
	go func() {
		defer close(errCh)
		fmt.Println("starting http server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			close(done)
			errCh <- fmt.Errorf("Server closed with error: %v", err)
			return
		}
		//log.Println("Server closed successfully")
	}()
	return errCh
}

func doneServer(ctx context.Context, srv *http.Server, done chan interface{}) chan error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	errCh := make(chan error)
	go func() {
		defer close(errCh)
		<-done
		if err := srv.Shutdown(ctx); err != nil {
			errCh <- fmt.Errorf("Failed to gracefully shutdown: %v", err)
			return
		}
		log.Println("Server shutdown gracefully")
	}()
	return errCh
}
