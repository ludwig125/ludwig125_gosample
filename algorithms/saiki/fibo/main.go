package main

import "fmt"

var num int

func main() {
	n := 7
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	fmt.Println("res:", fibo(n))
	fmt.Println("call:", num)
}

func fibo(i int) int {
	num++
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibo(i-1) + fibo(i-2)
}
