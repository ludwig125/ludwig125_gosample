package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i <= 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Blocked goroutiines %d\n", i)
			<-begin // チャネルから読み込めるようになるまでGoroutineは待機する
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutiines...")
	time.Sleep(3 * time.Second)
	close(begin) // chanのCloseですべてのGoroutineを解放する
	wg.Wait()
}
