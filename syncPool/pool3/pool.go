package main

import (
	"fmt"
	"sync"
)

// https://www.akshaydeo.com/blog/2017/12/23/How-did-I-improve-latency-by-700-percent-using-syncPool/

// Pool for our struct A
var pool *sync.Pool

// A dummy struct with a member
type A struct {
	Name string
}

// Func to init pool
func initPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("Returning new A")
			return new(A)
		},
	}
}

// Main func
func main() {
	// Initializing pool
	initPool()
	// Get hold of instance one
	one := pool.Get().(*A)
	one.Name = "first"
	fmt.Printf("one.Name = %s\n", one.Name)
	// Submit back the instance after using
	pool.Put(one)
	// Now the same instance becomes usable by another routine without allocating it again
}
