package main

import (
	"log"
	"time"
)

// Task はtask番号とtaskにかかるcost(作業時間)をまとめた構造体
type Task struct {
	Number int
	Cost   time.Duration
}

// taskを処理して処理済みのTask番号をSliceとして返す関数
func doTask(tasks []Task) []int {
	var doneTaskList []int
	for _, task := range tasks {
		log.Printf("do task number: %d\n", task.Number)
		// taskのための処理をする
		// ここではtask にかかるCostだけSleepする
		time.Sleep(task.Cost)
		doneTaskList = append(doneTaskList, task.Number) // 処理済みtask番号をlistにつめる
	}
	return doneTaskList
}

func main() {
	start := time.Now()

	// taskListに各Task番号とそのCostを定義する
	taskList := []Task{
		Task{1, 1 * time.Second},
		Task{2, 7 * time.Second},
		Task{3, 2 * time.Second},
		Task{4, 3 * time.Second},
		Task{5, 5 * time.Second},
		Task{6, 3 * time.Second},
	}

	count := 0
	for _, d := range doTask(taskList) { // 処理済みtaskの番号を読み出し
		count++
		log.Printf("done task number: %d\n", d)
	}
	log.Printf("Finished. Done %d tasks. Total time: %fs", count, time.Since(start).Seconds())
}
