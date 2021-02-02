package main

import (
	//	"fmt"
	"log"
	"time"
)

func main() {
	getNum := func() int {
		n := 2
		time.Sleep(1 * time.Second)
		return n
	}
	multipleNum := func(n int) int {
		mn := n * 3
		time.Sleep(1 * time.Second)
		return mn
	}
	log.Println("start getNum")
	n := getNum()
	log.Println("start multipleNum")
	m := multipleNum(n)
	log.Printf("end getNum %d\n", m)

	getChan := func() chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			ch <- 2
			time.Sleep(1 * time.Second)
		}()
		return ch
	}
	multipleChan := func(ch chan int) chan int {
		mch := make(chan int)
		go func() {
			defer close(mch)
			mch <- <-ch * 3
			time.Sleep(1 * time.Second)
		}()
		return mch
	}
	log.Println("start getChan")
	c := getChan()
	log.Println("start multipleChan")
	mc := multipleChan(c)
	log.Printf("end getNum %d\n", <-mc)
}
