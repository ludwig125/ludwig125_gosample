package main

import (
	"fmt"
	"sort"
)

func main() {
	k := []int{1, 3, 5, 7, 11, 12, 14, 15, 17, 19, 22, 24, 30, 41, 52, 54, 63, 66}
	fmt.Println(binarySearch(k, 54))
}

func binarySearch(k []int, x int) bool {
	sort.Ints(k)

	var l = 0
	var r = len(k)
	for {
		if (r - l) <= 0 {
			break
		}
		i := (r + l) / 2
		fmt.Println("n:", len(k), "i:", i, " k[i]:", k[i], "l, r", l, r)
		for j := l; j < r; j++ {
			fmt.Println("l ,r, k[j]:", l, r, k[j])
		}
		if k[i] == x {
			fmt.Println(k[i])
			return true
		} else if k[i] < x {
			fmt.Println("k[i] < x", l, i, i+1)
			l = i + 1
		} else {
			r = i
		}
	}
	return false
}
