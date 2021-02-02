package main

import "fmt"

var resA []int

func main() {
	n := 5
	// a := make([]int,n)
	// for i:=0;i<n;i++{
	// 	fmt.Scan(&a[i])
	// }
	a := []int{3, 5, 1, 2, 9}
	x := 8
	if solve(n, x, a) {
		fmt.Println("Yes", resA)
		return
	}
	fmt.Println("No")
}

func solve(i, x int, a []int) bool {
	//fmt.Printf("i: %d, sum: %d\n", i, sum)

	if i == 0 {
		return x == 0
	}
	if solve(i-1, x, a) {
		return true
	}
	if solve(i-1, x-a[i], a) {
		resA = append(resA, a[i])
		return true
	}
	return false
}
