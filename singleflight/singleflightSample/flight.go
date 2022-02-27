package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

// https://qiita.com/methane/items/27ccaee5b989fb5fca72

type Cache struct {
	m   sync.Mutex
	t   time.Time
	res interface{}
	err error
}

// DoDelay calls f() unless cached value which is created after calling time is available.
// If delay>0, sleep while it before calling f().  You can use it for sharing result
// value from more concurrent callers.
func (c *Cache) DoDelay(delay time.Duration, f func() (interface{}, error)) (v interface{}, err error) {
	t0 := time.Now()
	c.m.Lock()
	defer c.m.Unlock()

	// If c.t is newer, return cached value
	// We can't use `>=` because some system may produce exactly same time for multiple times.
	if c.t.Sub(t0) > 0 {
		return c.res, c.err
	}
	if delay > 0 {
		time.Sleep(delay)
	}

	c.t = time.Now()
	c.res, c.err = f()
	return c.res, c.err
}

// Do calls f() unless cached value which is created after calling time is available.
func (c *Cache) Do(f func() (interface{}, error)) (v interface{}, err error) {
	return c.DoDelay(0, f)
}

type NameRepo struct {
	m      sync.Mutex
	names  []string
	called int64 // GetNames() called
}

func (r *NameRepo) AddName(name string) {
	r.m.Lock()
	r.names = append(r.names, name)
	r.m.Unlock()
}

var GetNamesLatency = 1 * time.Millisecond

func (r *NameRepo) GetNames() map[string]bool {
	res := make(map[string]bool)

	r.m.Lock()
	r.called++
	for _, n := range r.names {
		res[n] = true
	}

	time.Sleep(GetNamesLatency) // レスポンスにかかる時間
	r.m.Unlock()
	return res
}

func Check(test string, nr *NameRepo, f func() map[string]bool) {
	start := time.Now()
	const N = 100
	var errorCount int64

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
					atomic.AddInt64(&errorCount, 1)
				}
			}
			wg.Done()
		}(name)
	}
	wg.Wait()

	fmt.Printf("[GetNamesLatency: %v] %s: called %d times, %d errors, %v duration\n", GetNamesLatency, test, nr.called, errorCount, time.Since(start))
}

func SampleDirect() {
	var nr NameRepo
	Check("direct", &nr, nr.GetNames)
}

func SampleSingleFlight() {
	var nr NameRepo
	var group singleflight.Group

	f := func() map[string]bool {
		v, _, _ := group.Do("", func() (interface{}, error) {
			return nr.GetNames(), nil
		})
		return v.(map[string]bool)
	}
	Check("singleflight", &nr, f)
}

func SampleZTC() {
	var nr NameRepo
	var cache Cache

	f := func() map[string]bool {
		v, _ := cache.Do(func() (interface{}, error) {
			return nr.GetNames(), nil
		})
		return v.(map[string]bool)
	}
	Check("ZeroTimeCache", &nr, f)
}

func SampleZTCDelay(delayTime time.Duration) {
	var nr NameRepo
	var cache Cache

	f := func() map[string]bool {
		v, _ := cache.DoDelay(delayTime, func() (interface{}, error) {
			return nr.GetNames(), nil
		})
		return v.(map[string]bool)
	}

	fmt.Println("sleep delay", delayTime)
	Check("ZeroTimeCacheDelay", &nr, f)
}

func main() {
	getNamesLatencys := []time.Duration{
		1 * time.Millisecond,
		5 * time.Millisecond,
		10 * time.Millisecond,
	}
	for _, getNamesLatency := range getNamesLatencys {
		GetNamesLatency = getNamesLatency

		SampleDirect()
		SampleSingleFlight()
		SampleZTC()

		ds := []time.Duration{
			10 * time.Microsecond,
			100 * time.Microsecond,
			1000 * time.Microsecond,
			5000 * time.Microsecond,
			10000 * time.Microsecond,
		}
		for _, d := range ds {
			SampleZTCDelay(d)
		}
	}
}
