package main

import (
	"fmt"
	"sync"
	"time"
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
	// Pass one to a go routine that just stops for 10 second
	// Assume that this could be any i/o bound task, something like an API call
	go func(o *A) {
		fmt.Printf("Before pool.Put\no.Name = %s\n", o.Name)
		time.Sleep(10 * time.Second)
		fmt.Printf("After pool.Put With same object\no.Name = %s\n", o.Name)
	}(one)
	// Main routine performs some operations for 1 second
	time.Sleep(1 * time.Second)
	// Till then your main routine submits back the main object
	pool.Put(one)
	two := pool.Get().(*A)
	two.Name = "second"
	fmt.Printf("two.Name = %s\n", two.Name)
	// Just for the sake of demo wait for next 15 seconds
	time.Sleep(15 * time.Second)
}
