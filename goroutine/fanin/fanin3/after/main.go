package main

import (
	"context"
	"log"
	"sync"
	"time"
)

// Task はtask番号とtaskにかかるcost(作業時間)をまとめた構造体
type Task struct {
	Number int
	Cost   time.Duration
}

// taskをchannel化するgenerator
func taskChannelGerenator(ctx context.Context, taskList []Task) <-chan Task {
	taskCh := make(chan Task)

	go func() {
		defer close(taskCh)
		for _, task := range taskList {
			select {
			case <-ctx.Done():
				return
			case taskCh <- task:
			}
		}
	}()
	return taskCh
}

// taskを処理して処理済みのTask番号をchannelとして返す関数
func doTask(ctx context.Context, taskCh <-chan Task) <-chan int {
	doneTaskCh := make(chan int)
	go func() {
		defer close(doneTaskCh)
		for task := range taskCh {
			select {
			case <-ctx.Done():
				return
			default:
				log.Printf("do task number: %d\n", task.Number)
				// taskのための処理をする
				// ここではtask にかかるCostだけSleepする
				time.Sleep(task.Cost)
				doneTaskCh <- task.Number // 処理済みtask番号をchannelにつめる
			}
		}
	}()
	return doneTaskCh
}

func merge(ctx context.Context, taskChs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	mergedTaskCh := make(chan int)

	mergeTask := func(taskCh <-chan int) {
		defer wg.Done()
		for t := range taskCh {
			select {
			case <-ctx.Done():
				return
			case mergedTaskCh <- t:
			}
		}
	}

	wg.Add(len(taskChs))
	for _, taskCh := range taskChs {
		go mergeTask(taskCh)
	}
	// 全てのtaskが処理されるまで待つ
	go func() {
		wg.Wait()
		close(mergedTaskCh)
	}()
	return mergedTaskCh
}

func main() {
	start := time.Now()
	// 処理の途中で中断されてもgoroutineリークしないようにcontextを使う(done channelでもいい)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// taskListに各Task番号とそのCostを定義する
	taskList := []Task{
		Task{1, 1 * time.Second},
		Task{2, 7 * time.Second},
		Task{3, 2 * time.Second},
		Task{4, 3 * time.Second},
		Task{5, 5 * time.Second},
		Task{6, 3 * time.Second},
	}

	taskCh := taskChannelGerenator(ctx, taskList)

	numWorkers := 4
	workers := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = doTask(ctx, taskCh)
	}

	count := 0
	for d := range merge(ctx, workers) { // mergeから処理済みtaskの番号を読み出し
		count++
		log.Printf("done task number: %d\n", d)
	}
	log.Printf("Finished. Done %d tasks. Total time: %fs", count, time.Since(start).Seconds())
}
