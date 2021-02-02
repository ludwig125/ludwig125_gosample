## 同時並列数の制御

#### 【同時並列数の制御】1. 並列数を制限しない場合

- 並列数を制限しない場合はこの通り単純
- 複数のgoroutineを起動する場合は、WaitGroupで待ち合わせをする
- ※time.Sleep(1 * time.Second)は処理の様子をわかりやすくするため入れているだけで、実用では必要ない
```
package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	doTask()
	log.Println("finished")
}

func doTask() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	var wg sync.WaitGroup
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fnA(n)

			// 処理をわかりやすくするため
			time.Sleep(1 * time.Second)
		}(num)
	}
	wg.Wait()
}

func fnA(n int) {
	log.Printf("do fnA. num: %d \n", n)
}
```
https://play.golang.org/p/JxdXBOThF0v

実行結果
```
2019/06/15 16:57:46 do fnA. num: 6 
2019/06/15 16:57:46 do fnA. num: 1 
2019/06/15 16:57:46 do fnA. num: 2 
2019/06/15 16:57:46 do fnA. num: 3 
2019/06/15 16:57:46 do fnA. num: 4 
2019/06/15 16:57:46 do fnA. num: 5 
2019/06/15 16:57:47 finished
```
- 全てのgoroutineが同時に起動して、それぞれ１秒Sleepしたあとでfinishedが出力されている

#### 【同時並列数の制御】2. 並列数を制限する場合

並列数を制限する場合
- 最大同時並列実行数をバッファサイズとしたチャネルを作り、そのチャネルの待ち合わせをすることで実現できる
- semチャネルは、一旦concurrency数だけ受信したらバッファがいっぱいになるので、「<-sem」が呼ばれて解放されない限り、後続のgoroutineは起動しない
=> 最大同時並列実行数を制限できる

```
package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	doTask()
	log.Println("finished")
}

const concurrency = 2 // 最大同時並列実行数

func doTask() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency) // concurrency数のバッファ
	for _, num := range numbers {
		sem <- struct{}{}

		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			defer func() { <-sem }() // 処理が終わったらチャネルを解放
			fnA(n)

			// 処理をわかりやすくするため
			time.Sleep(1 * time.Second)
		}(num)
	}
	wg.Wait()
}

func fnA(n int) {
	log.Printf("do fnA. num: %d \n", n)
}

```

https://play.golang.org/p/CEn0tw5SR-A

実行結果
```
2019/06/15 17:20:36 do fnA. num: 2 
2019/06/15 17:20:36 do fnA. num: 1 
2019/06/15 17:20:37 do fnA. num: 3 
2019/06/15 17:20:37 do fnA. num: 4 
2019/06/15 17:20:38 do fnA. num: 5 
2019/06/15 17:20:38 do fnA. num: 6 
2019/06/15 17:20:39 finished
```

- concurrency数ずつ（ここでは２つずつ）１秒おきに実行されていることがわかる

参考
- https://hori-ryota.com/blog/golang-channel-pattern/
- https://blog.monochromegane.com/blog/2015/12/15/how-to-speed-up-the-platinum-searcher-v2/
- https://qiita.com/kkohtaka/items/c42bfc75bede7cd8dc50
- https://gist.github.com/momotaro98/329ad3b039d5894f0f141090e957d4ad

上のコードの「sem <- struct{}{}」の後ろでlen(sem)を出力してみると、一旦semチャネルのバッファがconcurrency数=2に達したら、あとは２を保ったまま後続のgoroutineが起動しているのがわかる
```
sem <- struct{}{}
fmt.Printf("len(sem): %d\n", len(sem)) // <- バッファ内の値を出力
```

実行結果
```
len(sem): 1
len(sem): 2
2019/06/15 20:54:28 do fnA. num: 2 
2019/06/15 20:54:28 do fnA. num: 1 
len(sem): 2
2019/06/15 20:54:29 do fnA. num: 3 
len(sem): 2
2019/06/15 20:54:29 do fnA. num: 4 
len(sem): 2
2019/06/15 20:54:30 do fnA. num: 5 
len(sem): 2
2019/06/15 20:54:30 do fnA. num: 6 
2019/06/15 20:54:31 finished
```

#### 【同時並列数の制御】2-2. 並列数を制限する場合(チャネルを最後にcloseする)

上のをちょっと改良
- チャネルを使ったら最後にcloseしておいた方が安全なので、
- 以下のように全部のgoroutineを待って最後にチャネルをcloseするために、別のgoroutineを用意しておくと良い

- goroutineが１つだけの場合は最初のgo func()内に、「defer close(チャネル)」を呼び出せばいいが、今回のように複数のgoroutineを待つ場合はこのように書くのが良さそう
```
func doTask() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)
	for _, num := range numbers {
		sem <- struct{}{} // チャネルに送信

		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			defer func() { <-sem }()
			fnA(n)

			// 処理をわかりやすくするため
			time.Sleep(1 * time.Second)
		}(num)
	}

    // 別のgoroutineで上の全部のgoroutineが終わるまで待つ
    // 終わったらチャネルをclose
	go func() {
        defer close(sem)
		wg.Wait()
	}()
}
```

https://play.golang.org/p/0MbVqYjU-B3


#### 【同時並列数の制御】3. 並列数を制限してエラー処理もする場合

上のコードで、fnAがエラーを返す場合のエラー処理を入れる場合は以下になる

```
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	if err := doTask(); err != nil {
		log.Printf("error occured. %v", err)
	}
	log.Println("finished")
}

const concurrency = 2 // 最大同時並列実行数

var errFlag bool = true

func doTask() error {
	numbers := []int{1, 2, 3, 4, 5, 6}

	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)
	errChan := make(chan error, len(numbers))
	for _, num := range numbers {
		sem <- struct{}{} // チャネルに送信

		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			defer func() { <-sem }()
			if err := fnA(n); err != nil {
				errChan <- fmt.Errorf("failed to A, %v", err)
				log.Printf("--> fnA len(errChan) %d", len(errChan))

				time.Sleep(1 * time.Second) // 処理をわかりやすくするため
				return
			}
			time.Sleep(1 * time.Second) // 処理をわかりやすくするため
		}(num)
	}

	go func() {
		defer close(sem)
		defer close(errChan)
		wg.Wait()
	}()

	for err := range errChan {
		return err
	}
	return nil
}

func fnA(n int) error {
	log.Println("do fnA.")
	if errFlag {
		log.Printf("--> failed to do fnA. num: %d", n)
		return fmt.Errorf("error A. num: %d", n)
	}
	log.Printf("--> succeeded to do fnA. num: %d", n)
	return nil
}
```

- goroutine内で生じたエラーを外に伝えるために、errChanというチャネルを用意
`errChan := make(chan error, len(numbers))`
  - **このチャネルのバッファ数が重要！！**
- fnAの実行時にエラーが発生した場合はerrChanに送信
- errChanからエラーを読み取って、errを返す
```
for err := range errChan {
	return err
}
```
- wg.Wait()のあとに `close(errChan)` もする

errChanのバッファ数を起動されるgoroutineの数（ここではnumbersの6）だけ用意することで、エラーが複数発生してもチャネルが詰まらないようにしているのがポイント

https://play.golang.org/p/KdQB7fLn9Na

実行結果
```
2019/06/16 08:46:51 do fnA.
2019/06/16 08:46:51 --> failed to do fnA. num: 2
2019/06/16 08:46:51 --> fnA len(errChan) 1
2019/06/16 08:46:51 do fnA.
2019/06/16 08:46:51 --> failed to do fnA. num: 1
2019/06/16 08:46:51 --> fnA len(errChan) 2
2019/06/16 08:46:52 do fnA.
2019/06/16 08:46:52 --> failed to do fnA. num: 3
2019/06/16 08:46:52 --> fnA len(errChan) 3
2019/06/16 08:46:52 do fnA.
2019/06/16 08:46:52 --> failed to do fnA. num: 4
2019/06/16 08:46:52 --> fnA len(errChan) 4
2019/06/16 08:46:53 do fnA.
2019/06/16 08:46:53 --> failed to do fnA. num: 5
2019/06/16 08:46:53 --> fnA len(errChan) 5
2019/06/16 08:46:53 error occured. failed to A, error A. num: 2
2019/06/16 08:46:53 finished
```

- エラーが発生するたびにerrChanのバッファが埋まっていく様子がわかる

##### バッファ数が足りないとどうなるか？

試しに、errChanのバッファ数を０にすると、読み取り手がいないエラーを複数投げようとして詰まってdeadlockが発生する
`errChan := make(chan error)`

https://play.golang.org/p/Zy7xu6k8U9Y

実行結果
```
2019/06/16 08:25:37 do fnA.
2019/06/16 08:25:37 --> failed to do fnA.
2019/06/16 08:25:37 do fnA.
2019/06/16 08:25:37 --> failed to do fnA.
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
以下省略
```

参考
- https://christina04.hatenablog.com/entry/2016/06/22/022240

 ##### 上のコードが全部成功した場合
 一応載せておくとこんな感じ

 `var errFlag bool = false` にして実行する

https://play.golang.org/p/ZmRaPLpWC3T

実行結果
```
2019/06/16 08:18:39 do fnA.
2019/06/16 08:18:39 --> succeeded to do fnA. num: 2
2019/06/16 08:18:39 do fnA.
2019/06/16 08:18:39 --> succeeded to do fnA. num: 1
2019/06/16 08:18:40 do fnA.
2019/06/16 08:18:40 --> succeeded to do fnA. num: 3
2019/06/16 08:18:40 do fnA.
2019/06/16 08:18:40 --> succeeded to do fnA. num: 4
2019/06/16 08:18:41 do fnA.
2019/06/16 08:18:41 --> succeeded to do fnA. num: 5
2019/06/16 08:18:41 do fnA.
2019/06/16 08:18:41 --> succeeded to do fnA. num: 6
2019/06/16 08:18:42 finished
```

#### 【同時並列数の制御】4. contextを使ってエラー制御をきちんとする

上のエラーが起きたときの挙動を見てみると、エラーが起きてもすぐに終了していないことがわかる

上のコードで、起動時のnumを出力させて見ると以下のようになる
```
	for _, num := range numbers {
		sem <- struct{}{} // チャネルに送信
		log.Printf("num: %d", num)  ← 出力

		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			defer func() { <-sem }()
			log.Printf("goroutine num: %d", num) ← 出力
```

https://play.golang.org/p/mlzfPOWDDWt

実行結果
```
2019/06/16 16:51:31 num: 1
2019/06/16 16:51:31 num: 2
2019/06/16 16:51:31 goroutine n: 2
2019/06/16 16:51:31 do fnA.
2019/06/16 16:51:31 --> failed to do fnA. num: 2
2019/06/16 16:51:31 --> fnA len(errChan) 1
2019/06/16 16:51:31 goroutine n: 1
2019/06/16 16:51:31 do fnA.
2019/06/16 16:51:31 --> failed to do fnA. num: 1
2019/06/16 16:51:31 --> fnA len(errChan) 2
2019/06/16 16:51:32 num: 3
2019/06/16 16:51:32 goroutine n: 3
2019/06/16 16:51:32 do fnA.
2019/06/16 16:51:32 --> failed to do fnA. num: 3
2019/06/16 16:51:32 --> fnA len(errChan) 3
2019/06/16 16:51:32 num: 4
2019/06/16 16:51:32 goroutine n: 4
2019/06/16 16:51:32 do fnA.
2019/06/16 16:51:32 --> failed to do fnA. num: 4
2019/06/16 16:51:32 --> fnA len(errChan) 4
2019/06/16 16:51:33 num: 5
2019/06/16 16:51:33 goroutine n: 5
2019/06/16 16:51:33 do fnA.
2019/06/16 16:51:33 --> failed to do fnA. num: 5
2019/06/16 16:51:33 --> fnA len(errChan) 5
2019/06/16 16:51:33 num: 6
2019/06/16 16:51:33 error occured. failed to A, error A. num: 2
2019/06/16 16:51:33 finished
```

これはリソースの無駄なので、エラーが起きたら即終了させるようにしたい

こういうときはcontextが便利


「【同時並列数の制御】3」のソースコードをcontextを使って以下のように書き直す

```
package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	if err := doTask(); err != nil {
		log.Printf("error occured. %v", err)
	}
	log.Println("finished")
}

const concurrency = 2 // 最大同時並列実行数

var errFlag bool = true

func doTask() error {
	numbers := []int{1, 2, 3, 4, 5, 6}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background()) // contextとキャンセル関数を定義
	defer cancel() // doTask終了時に子プロセスを全て終了するようにしたい

	sem := make(chan struct{}, concurrency)
	errChan := make(chan error, len(numbers))
	for _, num := range numbers {
		sem <- struct{}{} // チャネルに送信
		log.Printf("num: %d", num)

		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			defer func() { <-sem }()

			log.Printf("goroutine num: %d", n)
			select {
			case <-ctx.Done(): // contextのcancelが呼び出されたらここに入って即終了する
				return
			default:
			}
			if err := fnA(n); err != nil {
				errChan <- fmt.Errorf("failed to A, %v", err)
				log.Printf("--> fnA len(errChan) %d", len(errChan))

				// エラーが発生したら他の処理はキャンセル
				cancel()
				time.Sleep(1 * time.Second) // 処理をわかりやすくするため
				return
			}
			time.Sleep(1 * time.Second) // 処理をわかりやすくするため
		}(num)
	}

	go func() {
		defer close(sem)
		defer close(errChan)
		wg.Wait()
	}()

	for err := range errChan {
		return err
	}
	return nil
}

func fnA(n int) error {
	log.Println("do fnA.")
	if errFlag {
		log.Printf("--> failed to do fnA. num: %d", n)
		return fmt.Errorf("error A. num: %d", n)
	}
	log.Printf("--> succeeded to do fnA. num: %d", n)
	return nil
}
```
-  contextのcancelが呼び出されたら「<-ctx.Done()」を受け取って即終了するようにする
```
select {
case <-ctx.Done():
	return
default:
}
```
- エラーが発生したら他の処理はキャンセルするため `cancel()` を送る

https://play.golang.org/p/N1mjZlo51VV

実行結果
```
2019/06/16 16:54:11 num: 1
2019/06/16 16:54:11 num: 2
2019/06/16 16:54:11 goroutine num: 3
2019/06/16 16:54:11 do fnA.
2019/06/16 16:54:11 --> failed to do fnA. num: 2
2019/06/16 16:54:11 --> fnA len(errChan) 1
2019/06/16 16:54:11 goroutine num: 3
2019/06/16 16:54:11 num: 3
2019/06/16 16:54:11 goroutine num: 4
2019/06/16 16:54:11 num: 4
2019/06/16 16:54:11 goroutine num: 5
2019/06/16 16:54:11 num: 5
2019/06/16 16:54:11 goroutine num: 6
2019/06/16 16:54:11 num: 6
2019/06/16 16:54:11 error occured. failed to A, error A. num: 2
2019/06/16 16:54:11 finished
```

- 「do fnA. 」は一度しか呼び出されていない
- 一つエラーが発生したら、それ以外のgoroutineは起動してもすぐに処理が終わっていることがわかる

#### 【同時並列数の制御】5. contextに加えてerrgroupを使ってエラー制御をかんたんにする

errgroupを使うことで、エラー制御が便利になる。

以下は、syncの代わりにerrgroupを使っている

- `go get golang.org/x/sync/errgroup` でerrgroupを取得
- errChanは使わないで済むようになった

```
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := doTask(); err != nil {
		log.Printf("error occured. %v", err)
	}
	log.Println("finished")
}

const concurrency = 2 // 最大同時並列実行数

var errFlag bool = true

func doTask() error {
	numbers := []int{1, 2, 3, 4, 5, 6}

	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sem := make(chan struct{}, concurrency)
	for _, num := range numbers {
		sem <- struct{}{} // チャネルに送信
		log.Printf("num: %d", num)

		n := num
		eg.Go(func() error {
			defer func() { <-sem }()

			log.Printf("goroutine num: %d", n)
			select {
			case <-ctx.Done():
				//return ctx.Err()
				return nil
			default:
			}
			if err := fnA(n); err != nil {
				// エラーが発生したら他の処理はキャンセル
				cancel()
				time.Sleep(1 * time.Second) // 処理をわかりやすくするため
				return fmt.Errorf("failed to A, %v", err)
			}
			time.Sleep(1 * time.Second) // 処理をわかりやすくするため
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		defer close(sem)
		return err
	}
	return nil
}

func fnA(n int) error {
	log.Println("do fnA.")
	if errFlag {
		log.Printf("--> failed to do fnA. num: %d", n)
		return fmt.Errorf("error A. num: %d", n)
	}
	log.Printf("--> succeeded to do fnA. num: %d", n)
	return nil
}
```

https://play.golang.org/p/Bk_nR55k_Ng

実行結果
```
2019/06/16 17:17:31 num: 1
2019/06/16 17:17:31 num: 2
2019/06/16 17:17:31 goroutine num: 2
2019/06/16 17:17:31 do fnA.
2019/06/16 17:17:31 --> failed to do fnA. num: 2
2019/06/16 17:17:31 goroutine num: 1
2019/06/16 17:17:31 num: 3
2019/06/16 17:17:31 goroutine num: 3
2019/06/16 17:17:31 num: 4
2019/06/16 17:17:31 goroutine num: 4
2019/06/16 17:17:31 num: 5
2019/06/16 17:17:31 goroutine num: 5
2019/06/16 17:17:31 num: 6
2019/06/16 17:17:31 goroutine num: 6
2019/06/16 17:17:32 error occured. failed to A, error A. num: 2
2019/06/16 17:17:32 finished
```

参考：
- http://dono.hatenablog.com/entry/2018/01/04/111204
- https://note.mu/kltl/n/na70c3eec41ca
- https://tomokazu-kozuma.com/how-to-use-sync-waitgropu-and-errorgroup-group-to-summarize-parallel-processing-with-golang/
- https://www.oreilly.com/learning/run-strikingly-fast-parallel-file-searches-in-go-with-sync-errgroup
- https://deeeet.com/writing/2016/10/12/errgroup/
- https://godoc.org/golang.org/x/sync/errgroup
