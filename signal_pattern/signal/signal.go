package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	defer fmt.Println("done")

	tranpSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT}

	// 受信するチャネル
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, tranpSignals...)

	go doMain()

	sig := <-sigCh
	fmt.Println("Got signal!", sig)
}

func doMain() {
	defer fmt.Println("done infinite loop")
	for {
		fmt.Println("sleep 1")
		time.Sleep(1 * time.Second)
	}
}
