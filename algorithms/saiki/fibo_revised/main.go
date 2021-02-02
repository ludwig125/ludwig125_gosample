package main

import "fmt"

var num int

func main() {
	n := 7
	memo := make([]int, n+2)
	for i := 2; i <= n+1; i++ {
		memo[i] = -1
	}
	fmt.Println("res:", fibo(n, &memo))
	fmt.Println("call:", num)
}

func fibo(i int, memo *([]int)) int {
	num++
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}

	if (*memo)[i] >= 0 {
		//fmt.Println("koko i:", i, (*memo), (*memo)[i])
		return (*memo)[i]
	}
	(*memo)[i] = fibo(i-1, memo) + fibo(i-2, memo)
	return (*memo)[i]
}

// var num int

// func main() {
// 	n := 7
// 	memo := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		memo[i] = -1
// 	}
// 	fmt.Println("res:", fibo(n, &memo))
// 	fmt.Println("call:", num)
// }

// func fibo(i int, memo *([]int)) int {
// 	num++
// 	if i == 0 {
// 		return 0
// 	}
// 	if i == 1 {
// 		return 1
// 	}

// 	if (*memo)[i] >= 0 {
// 		//fmt.Println("koko i:", i, (*memo), (*memo)[i])
// 		return (*memo)[i]
// 	}
// 	(*memo)[i] = fibo(i-1, memo) + fibo(i-2, memo)
// 	return (*memo)[i]
// }
