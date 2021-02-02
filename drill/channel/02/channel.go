package main

import (
    "log"
    "time"
)

func f(ch chan bool) {
    time.Sleep(1 * time.Second)
    ch <- true
}

func main() {
    ch := make(chan bool)
    go f(ch)
    log.Println(<-ch)
}
