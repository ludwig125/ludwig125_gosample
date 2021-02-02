package main

import (
	"context"
	"log"
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
			case taskCh <- task: // taskをchannelにつめる
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
			}
			log.Printf("do task number: %d\n", task.Number)
			// taskのための処理をする
			// ここではtask にかかるCostだけSleepする
			time.Sleep(task.Cost)
			doneTaskCh <- task.Number // 処理済みtask番号をchannelにつめる
		}
	}()
	return doneTaskCh
}

func main() {
	start := time.Now()
	// 処理の途中で中断されてもgoroutineリークしないようにcontextを使う
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

	// taskChannelGerenatorとdoTaskという２つのステージをまとめたpipelineを定義する
	pipeline := doTask(ctx, taskChannelGerenator(ctx, taskList))
	count := 0
	for d := range pipeline { // pipelineから処理済みtaskの番号を読み出し
		count++
		log.Printf("done task number: %d\n", d)
	}

	log.Printf("Finished. Done %d tasks. Total time: %fs", count, time.Since(start).Seconds())
}
