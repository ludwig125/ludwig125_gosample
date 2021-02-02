// https://github.com/kat-co/concurrency-in-go-src/blob/4e55fd7f3f5b9c5efc45a841702393a1485ba206/concurrency-patterns-in-go/confinement/fig-confinement-structs.go
package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3]) // dataの中の先頭の３バイトを含んだスライスを渡す
	go printData(&wg, data[3:]) // dataの中の後半の３バイトを含んだスライスを渡す

	wg.Wait()
}
