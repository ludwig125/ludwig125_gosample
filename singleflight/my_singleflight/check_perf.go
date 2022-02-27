package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

// https://qiita.com/methane/items/27ccaee5b989fb5fca72

type GetResponser struct {
	m      sync.Mutex
	names  []string
	called int64 // GetResponse() called
}

func (r *GetResponser) AddName(name string) {
	r.m.Lock()
	r.names = append(r.names, name)
	r.m.Unlock()
}

var GetResponseLatency = 1 * time.Millisecond

func (r *GetResponser) GetResponse() map[string]bool {
	res := make(map[string]bool)

	r.m.Lock()
	r.called++
	for _, n := range r.names {
		res[n] = true
	}

	time.Sleep(GetResponseLatency) // レスポンスにかかる時間
	r.m.Unlock()
	return res
}

func DoDirect() {
	var nr GetResponser
	Check("direct", &nr, nr.GetResponse)
}

func DoSingleFlight() {
	var nr GetResponser
	var group singleflight.Group

	f := func() map[string]bool {
		v, _, _ := group.Do("", func() (interface{}, error) {
			return nr.GetResponse(), nil
		})
		return v.(map[string]bool)
	}
	Check("singleflight", &nr, f)
}

func DoMySingleFlight() {
	var nr GetResponser
	flight := NewFlight()

	f := func() map[string]bool {
		v, _, _ := group.Do("", func() (interface{}, error) {
			return nr.GetResponse(), nil
		})
		return v.(map[string]bool)
	}
	Check("singleflight", &nr, f)
}

func Check(test string, nr *GetResponser, f func() map[string]bool) {
	start := time.Now()
	const N = 100
	var notFoundCount int64

	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		name := fmt.Sprintf("worker-%d", i)
		go func(name string) {
			for j := 0; j < 10; j++ {
				n := fmt.Sprintf("%s-%d", name, j)
				nr.AddName(n)
				result := f()
				if !result[n] {
					atomic.AddInt64(&notFoundCount, 1)
				}
			}
			wg.Done()
		}(name)
	}
	wg.Wait()

	fmt.Printf("[GetResponseLatency: %v] %s: called %d, not found: %d, duration: %v\n", GetResponseLatency, test, nr.called, notFoundCount, time.Since(start))
}

func main() {
	getNamesLatencys := []time.Duration{
		1 * time.Millisecond,
		// 5 * time.Millisecond,
		// 10 * time.Millisecond,
	}
	for _, getNamesLatency := range getNamesLatencys {
		GetResponseLatency = getNamesLatency

		DoDirect()
		DoSingleFlight()
		// SampleZTC()

		// ds := []time.Duration{
		// 	10 * time.Microsecond,
		// 	100 * time.Microsecond,
		// 	1000 * time.Microsecond,
		// 	5000 * time.Microsecond,
		// 	10000 * time.Microsecond,
		// }
		// for _, d := range ds {
		// 	SampleZTCDelay(d)
		// }
	}
}
