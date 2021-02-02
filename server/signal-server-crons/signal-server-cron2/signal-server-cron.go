package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type batchCtl struct {
	Mu     sync.Mutex
	Status int
}

func (b *batchCtl) runBatch(ctx context.Context) error {
	log.Println("start batch")
	defer log.Println("end batch")
	if err := createFile(); err != nil {
		return err
	}
	for i := 0; i < 100; i++ {
		select {
		case <-ctx.Done(): // キャンセルが発生した場合
			log.Printf("! cancel done %v\n", ctx.Err())
			return ctx.Err()
		default:
		}
		if err := writeFile(fmt.Sprintf("%d\n", i)); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (b *batchCtl) getStatus() int {
	return b.Status
}
func (b *batchCtl) updateStatus(i int) {
	b.Mu.Lock()
	defer b.Mu.Unlock()
	b.Status = i
}

var status int

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	defer log.Println("finished main")

	codeCh := make(chan string)
	mux := http.NewServeMux()
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "status: %d\n", status)
	})
	mux.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Start process"))
		// ctx, cancel := context.WithCancel(ctx)
		// defer cancel()
		// batchFn(ctx)
	})
	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		log.Println("received stop")
		w.Write([]byte("Stop process"))
		// Cancel the context on request
		// cancel()
		codeCh <- "stop"
		close(codeCh)
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			log.Fatalln("Server closed with error:", err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errChan := make(chan error, 1)
	go func() {
		errChan <- batchFn(ctx)
	}()
	select {
	case code := <-codeCh:
		// Post process after shutdown here
		//s.Shutdown(context.Background())
		log.Printf("Got code=%s", code)
	case err := <-errChan:
		log.Printf("errChan err %v", err)
	default:
	}

	// c := cron.New()
	// c.AddFunc("@every 20s", func() {
	// 	b := batchCtl{}

	// 	errChan := make(chan error, 1)
	// 	go func() {
	// 		defer func() { status = b.getStatus() }()
	// 		defer b.updateStatus(0)
	// 		b.updateStatus(1)
	// 		status = b.getStatus()
	// 		errChan <- b.runBatch(ctx)
	// 	}()

	// 	select {
	// 	case <-ctx.Done(): // キャンセルが発生した場合
	// 		log.Printf("! cancel done %v\n", ctx.Err())
	// 	case err := <-errChan: // heavyFuncの結果を取得した場合
	// 		log.Printf("err %v\n", err)
	// 	}

	// })
	// c.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		log.Println("Failed to gracefully shutdown:", err)
	}
	log.Println("Server shutdown")
}

// func doBatch(ctx context.Context) {
// 	ctx, cancel := context.WithTimeout(ctx, 10*time.Second) // 全体2秒でタイムアウト
// 	defer cancel()
// 	batchFn(ctx)
// }
func batchFn(ctx context.Context) error {
	b := batchCtl{}

	defer func() { status = b.getStatus() }()
	defer b.updateStatus(0)
	b.updateStatus(1)
	status = b.getStatus()
	return b.runBatch(ctx)

	// errChan := make(chan error, 1)
	// go func() {
	// 	defer func() { status = b.getStatus() }()
	// 	defer b.updateStatus(0)
	// 	b.updateStatus(1)
	// 	status = b.getStatus()
	// 	errChan <- b.runBatch(ctx)
	// }()

	// select {
	// case <-ctx.Done(): // キャンセルが発生した場合
	// 	log.Printf("! parent cancel done %v\n", ctx.Err())
	// 	return
	// case err := <-errChan: // heavyFuncの結果を取得した場合
	// 	log.Printf("err %v\n", err)
	// 	return
	// }

}
