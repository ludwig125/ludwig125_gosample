package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
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

// func getNumber() {
// 	read := make(chan int)
// 	write := make(chan int)

// 	writeStop := make(chan struct{})
// 	readStop := make(chan struct{})
// 	doneCh := make(chan struct{})

// 	// Create an instance of `SafeNumber`
// 	i := &SafeNumber{}

// 	n := 3
// 	// Use `Set` and `Get` instead of regular assignments and reads
// 	// We can now be sure that we can read only if the write has completed, or vice versa
// 	var wg sync.WaitGroup
// 	for j := 0; j < n; j++ {
// 		wg.Add(1)
// 		go func(j int) {
// 			defer wg.Done()
// 			<-read
// 			fmt.Println("set", j)
// 			i.Set(j)
// 			write <- 1
// 		}(j)
// 	}

// 	go func() {
// 		cnt := 0
// 		for {
// 			select {
// 			case <-readStop:
// 				fmt.Println("read stop!")
// 				close(doneCh)
// 				return
// 			case <-write:
// 				fmt.Println("get", i.Get())
// 				cnt++
// 			read <- 1
// 		}
// 	}
// }()

// read <- 1
// // time.Sleep(1 * time.Second)
// wg.Wait()
// // close(stopCh)
// // close(write)
// //close(read)

// <-doneCh
// }

// Set とGetはバラバラに動く
func getNumber() {
	rand.Seed(time.Now().UnixNano())

	read := make(chan int)
	write := make(chan int)

	writeStop := make(chan struct{})
	//readStop := make(chan struct{})
	//doneCh := make(chan struct{})

	// Create an instance of `SafeNumber`
	i := &SafeNumber{}

	n := 10
	// Use `Set` and `Get` instead of regular assignments and reads
	// We can now be sure that we can read only if the write has completed, or vice versa
	var wg sync.WaitGroup
	for j := 0; j < n; j++ {
		go func(j int) {
			// defer wg.Done()
			for {
				select {
				case <-writeStop:
					fmt.Println("write stop!")
					//fmt.Println("call read stop!")
					//close(readStop)
					return
				case <-read:
					sleepTime := rand.Intn(100)
					fmt.Printf("random number is %v Millisecond\n", sleepTime)
					time.Sleep(time.Duration(sleepTime) * time.Millisecond)
					fmt.Println("set", j)
					i.Set(j)
					write <- 1
				}
			}
		}(j)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		cnt := 0
		for {
			select {
			// case <-readStop:
			// 	fmt.Println("read stop!")
			// 	close(doneCh)
			// 	return
			case <-write:
				fmt.Println("get", i.Get())
				cnt++
				if cnt == n {
					fmt.Println("call write stop!")
					close(writeStop)
					return
				}
				read <- 1
			}
		}
	}()

	read <- 1
	// time.Sleep(1 * time.Second)
	wg.Wait()
	// close(stopCh)
	// close(write)
	//close(read)

	//<-doneCh
}
