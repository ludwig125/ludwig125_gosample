package main

import (
	"fmt"
	"sync"
)

func main() {
	getNumber()
}

// First, create a struct that contains the value we want to return
// along with a mutex instance
type SafeNumber struct {
	val int
	m   sync.Mutex
}

func (i *SafeNumber) Get() int {
	// The `Lock` method of the mutex blocks if it is already locked
	// if not, then it blocks other calls until the `Unlock` method is called
	i.m.Lock()
	// Defer `Unlock` until this method returns
	defer i.m.Unlock()
	// Return the value
	return i.val
}

func (i *SafeNumber) Set(val int) {
	// Similar to the `Get` method, except we Lock until we are done
	// writing to `i.val`
	i.m.Lock()
	defer i.m.Unlock()
	i.val = val
}

// Set とGetはバラバラに動く
func getNumber() {
	stopCh := make(chan struct{})
	doneCh := make(chan struct{})
	//readCh := make(chan struct{})

	// Create an instance of `SafeNumber`
	i := &SafeNumber{}
	// Use `Set` and `Get` instead of regular assignments and reads
	// We can now be sure that we can read only if the write has completed, or vice versa
	var wg sync.WaitGroup
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			fmt.Println("set", j)
			i.Set(j)
			//readCh <- struct{}{}
		}(j)
	}

	go func() {
		for {
			select {
			case <-stopCh:
				fmt.Println("stop!")
				close(doneCh)
				return
			default:
			}
			fmt.Println("get", i.Get())
		}
	}()
	// time.Sleep(1 * time.Second)
	wg.Wait()
	close(stopCh)

	<-doneCh
}
