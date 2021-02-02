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
	if solve(n, x, a, 0, 0) {
		fmt.Println("Yes", resA)
		return
	}
	fmt.Println("No")
}

func solve(n, x int, a []int, i, sum int) bool {
	fmt.Printf("i: %d, sum: %d\n", i, sum)

	if i == n {
		return sum == x
	}
	if solve(n, x, a, i+1, sum) {
		return true
	}
	if solve(n, x, a, i+1, sum+a[i]) {
		resA = append(resA, a[i])
		return true
	}
	return false
}
