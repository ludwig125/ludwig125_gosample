package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

// ref: https://stackoverflow.com/questions/57166480/golang-http-server-graceful-shutdown-after-a-signal

var (
    simpleHTTPServer      http.Server
    sigChan               chan os.Signal
    simpleServiceShutdown chan bool
)

func hello(w http.ResponseWriter, r *http.Request) {
	// log.Println("heavy process starts")
	// time.Sleep(5 * time.Second)
	// log.Println("done")
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("hello\n"))
}

func simpleServiceStarter() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
    simpleHTTPServer := &http.Server{
        Addr:           ":9000",
        Handler:        mux,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    fmt.Printf("\nstarting http server on :9000 ")
    err := simpleHTTPServer.ListenAndServe()
    if err != http.ErrServerClosed {
        fmt.Printf("error starting simple service or closing listener - %v\n", err)
    }
    fmt.Printf("simple service http server shutdown completed - %v\n", err)

	// communicate with main thread
	fmt.Printf("send simpleServiceShutdown")
    simpleServiceShutdown <- true
}

func signalHandler() {
    // Handle SIGINT and SIGHUP.
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGHUP)

    sig := <-sigChan
    fmt.Printf("\nsignalHandler() received signal: %v\n", sig)

    // gracefully shutdown http server
    err := simpleHTTPServer.Shutdown(context.Background())
    fmt.Printf("simple service shutdown on signal %v, error: %v\n", sig, err)
    close(sigChan)
}

func main() {
    // block all async signals to this server. And we register only SIGINT and SIGHUP for now.
    signal.Ignore()

    sigChan = make(chan os.Signal, 1)
    simpleServiceShutdown = make(chan bool)
    go signalHandler()

    go simpleServiceStarter()
	<-simpleServiceShutdown // wait server to shutdown
	fmt.Printf("receive simpleServiceShutdown")
    close(simpleServiceShutdown)
}