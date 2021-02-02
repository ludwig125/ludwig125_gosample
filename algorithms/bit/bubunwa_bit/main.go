package main

import "fmt"

func main() {
	n := 3
	a := []int{7, 2, 9}
	k := 11
	for bit := 0; bit < 1<<n; bit++ {
		sum := 0
		for i := 0; i < n; i++ {
			if (bit & (1 << i)) > 0 { // bitのi番目に1が立っているか？
				sum += a[i]
			}
		}
		if k == sum {
			fmt.Println(bit, sum)
		}
	}
}
