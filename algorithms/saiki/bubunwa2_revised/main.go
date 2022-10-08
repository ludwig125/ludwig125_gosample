package main

import "fmt"

func main() {
	fmt.Println(fibo(5))
}

func fibo(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}

	res := fibo(i-2) + fibo(i-1)
	fmt.Println(res)
	return res
}
