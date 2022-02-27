package fetcher

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

// ErrExecutorNotFound is used when query executor was not found for given user type.
var ErrExecutorNotFound = errors.New("executor not found")

// Fetcher executes query and cache result.
type Fetcher struct {
	m map[int8]*executor
}

// Fetcher executes query and cache result.
type executor struct {
	mu      sync.Mutex
	m       map[string]*call
	doQuery Query
}

type call struct {
	wg sync.WaitGroup

	// These fields are written once before the WaitGroup is done
	// and are only read after the WaitGroup is done.
	val map[string]interface{}
	err error
}

// Query is actual query executor.
type Query func(Request) Response

// Request is reqeust of Query
type Request struct {
	Ctx      context.Context
	UserID   string
	UserType int8
	key      string
	c        chan Response
}

// Response is response of Query
type Response struct {
	Result map[string]interface{}
	Err    error
}

// StartTime is.
var StartTime time.Time // TODO: remove in future

// New creates new Fetcher and start it.
func New(doQuery Query) *Fetcher {
	// ロック競合を緩和するため、user typeと同じ数だけのexecutorを作る
	//
	// rtg_beams_client.goでも独自にuser typeを定義している
	// entityのものと同じはずだが、ずれないように注意する
	m := make(map[int8]*executor)
	for _, userType := range []int8{0, 1, 2, 4, 10, 11, 12} {
		m[userType] = &executor{
			m:       make(map[string]*call),
			doQuery: doQuery,
		}
	}

	f := &Fetcher{
		m: m,
	}

	StartTime = time.Now() // TODO: remove in future
	log.Println("fetcher StartTime", StartTime)

	return f
}

// Do executes Query and returns its result.
func (f *Fetcher) Do(ctx context.Context, userID string, userType int8) (map[string]interface{}, error) {
	e, found := f.m[userType]
	if !found {
		// executor見つからなければエラーを返して終了する
		// ここでexecutorを追加するようにすると`f.m`をロックする必要が生じ、
		// 全ての処理に対してロック競合が発生するので、
		// user typeごとにexecutorを用意した意味がなくなってしまう
		return nil, fmt.Errorf("%w: userType=%d", ErrExecutorNotFound, userType)
	}
	return e.do(ctx, userID, userType)
}

func (e *executor) do(ctx context.Context, userID string, userType int8) (map[string]interface{}, error) {
	requestID, _ := requestid.RequestID(ctx)
	if len(requestID) == 0 {
		requestID = requestid.Generate()
	}
	// requestID + userIDが同じリクエストでクエリの結果を共有する
	// TODO: リトライ回数も入れた方が安全
	// リトライ間でデータを共有すると、失敗したリトライの結果(空)を参照し続けることになる
	// 現在はリトライ側でリクエストIDを作り直すか、あえてリクエストIDをセットせずfetcherに作らせることでリトライ間のデータ共有を回避している
	key := fmt.Sprintf("%s:%s:%s", requestID, userID, strconv.FormatInt(int64(userType), 10))
	req := Request{ctx, userID, userType, key, make(chan Response, 1)}

	e.mu.Lock()
	// e.mは必ず初期化されている(fetcher.Doのコメントを参照)
	if c, ok := e.m[req.key]; ok {
		// 2番目以降にfetcher.Doを呼ぶとこの分岐に入り、クエリの結果がセットされるのを待ち続ける
		e.mu.Unlock()
		// チャネルを使えばselectでタイムアウトできるようになるが、性能が劣化したので不採用
		c.wg.Wait()

		return c.val, c.err
	}
	c := new(call)
	c.wg.Add(1)
	e.m[req.key] = c
	e.mu.Unlock()

	res := e.doQuery(req)
	c.val, c.err = res.Result, res.Err

	go func() {
		// <-ctx.Done()にすると、fetcher.Doを呼んだユースケースが終了した時点でキャンセルが発生してしまう
		// 汎用性を高めるなら、スリープ時間を設定値にしてもよい
		// time.Sleep(50 * time.Millisecond)
		dynamicSleep() // TODO: お試し

		e.mu.Lock()
		delete(e.m, req.key)
		e.mu.Unlock()
	}()

	c.wg.Done()
	return c.val, c.err
}

var count int

func dynamicSleep() {
	t := time.Now()
	if t.Sub(StartTime) < 10*60*time.Second {
		checkCountAndPrint(t, StartTime, 0)
		time.Sleep(100 * time.Microsecond)
		return
	}
	if t.Sub(StartTime) < 20*60*time.Second {
		checkCountAndPrint(t, StartTime, 1)
		time.Sleep(500 * time.Microsecond)
		return
	}
	if t.Sub(StartTime) < 30*60*time.Second {
		checkCountAndPrint(t, StartTime, 2)
		time.Sleep(1 * time.Millisecond)
		return
	}
	if t.Sub(StartTime) < 40*60*time.Second {
		checkCountAndPrint(t, StartTime, 3)
		time.Sleep(5 * time.Millisecond)
		return
	}
	if t.Sub(StartTime) < 50*60*time.Second {
		checkCountAndPrint(t, StartTime, 4)
		time.Sleep(10 * time.Millisecond)
		return
	}
	if t.Sub(StartTime) < 60*60*time.Second {
		checkCountAndPrint(t, StartTime, 5)
		time.Sleep(50 * time.Millisecond)
		return
	}
	if t.Sub(StartTime) < 70*60*time.Second {
		checkCountAndPrint(t, StartTime, 6)
		time.Sleep(100 * time.Millisecond)
		return
	}
	if t.Sub(StartTime) >= 70*60*time.Second {
		checkCountAndPrint(t, StartTime, 7)
		// sleepなし
		return
	}
}

//　大量のログを出したくないので、各レンジに付き１回ずつだけ出力する
func checkCountAndPrint(t, start time.Time, compare int) {
	if count == compare {
		count++
		log.Printf("[checkCountAndPrint] now %v, start %v", t, start)
	}
}
