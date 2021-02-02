package main

import (
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

func (b *batchCtl) runBatch() error {
	log.Println("start batch")
	if err := createFile(); err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		if err := writeFile(fmt.Sprintf("%d\n", i)); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	log.Println("end batch")

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
	defer fmt.Println("done")

	tranpSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT}

	// 受信するチャネル
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, tranpSignals...)

	http.HandleFunc("/api", apiHandler)
	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	var wg sync.WaitGroup
	b := batchCtl{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() { status = b.getStatus() }()
		defer b.updateStatus(0)
		b.updateStatus(1)
		status = b.getStatus()
		b.runBatch()
	}()
	// fmt.Println(b.Status)
	// time.Sleep(3 * time.Second)
	// fmt.Println(b.Status)
	wg.Wait()

	sig := <-sigCh
	fmt.Println("Got signal!", sig)
	// time.Sleep(10 * time.Second)
	//fmt.Println(b.Status)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status: %d\n", status)

	// クエリパラメータ取得してみる
	fmt.Fprintf(w, "クエリ：%s\n", r.URL.RawQuery)

	queryValues := r.URL.Query()
	fmt.Fprintf(w, "mode %s status\n", queryValues.Get("mode"))

	// Bodyデータを扱う場合には、事前にパースを行う
	r.ParseForm()

	// Formデータを取得.
	form := r.PostForm
	fmt.Fprintf(w, "フォーム：\n%v\n", form)

	// または、クエリパラメータも含めて全部.
	params := r.Form
	fmt.Fprintf(w, "フォーム2：\n%v\n", params)
}

// func main() {
// 	defer fmt.Println("done")

// 	tranpSignals := []os.Signal{
// 		syscall.SIGHUP,
// 		syscall.SIGINT,
// 		syscall.SIGTERM,
// 		syscall.SIGQUIT}

// 	// 受信するチャネル
// 	sigCh := make(chan os.Signal, 1)

// 	signal.Notify(sigCh, tranpSignals...)

// 	go doMain()

// 	sig := <-sigCh
// 	fmt.Println("Got signal", sig)
// }

// func doMain() {
// 	defer fmt.Println("done infinite loop")
// 	for {
// 		fmt.Println("sleep 1")
// 		time.Sleep(1 * time.Second)
// 	}
// }
