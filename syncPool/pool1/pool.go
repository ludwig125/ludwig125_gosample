package main

import (
	"log"
	"sync"
	"time"
)

// https://qiita.com/ReSTARTR/items/ee943512243aedb3aa25

type worker struct{}

func (w *worker) work() {
	// time.Sleep(time.Duration(500+rand.Intn(100)) * time.Millisecond)
	time.Sleep(1000 * time.Millisecond)
}

type workers struct {
	limit chan struct{}
	pool  sync.Pool
}

func newWorkers(n int) *workers {
	ws := workers{}
	ws.limit = make(chan struct{}, n)
	ws.pool = sync.Pool{New: func() interface{} {
		return &worker{}
	}}
	return &ws
}

func (ws *workers) Get() *worker {
	ws.limit <- struct{}{}
	return ws.pool.Get().(*worker)
}

func (ws *workers) Put(w *worker) {
	ws.pool.Put(w)
	<-ws.limit
}

func main() {
	ws := newWorkers(5)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			w := ws.Get()
			log.Printf("[worker %v] --> goroutine %d", w, i)
			w.work()
			log.Printf("[worker %v] <-- goroutine %d", w, i)
			ws.Put(w)
		}(i)
	}
	wg.Wait()
}
