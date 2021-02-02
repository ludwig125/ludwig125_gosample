package main

import "fmt"

func main() {
	// bit := 11
	// fmt.Println(bit & (1 << 3))
	n := 5
	for bit := 0; bit < 1<<n; bit++ {

		var s []int
		for i := 0; i < n; i++ {
			if (bit & (1 << i)) > 0 { // bitのi番目に1が立っているか？
				s = append(s, i)
			}
		}
		fmt.Println(bit, s)

		// fmt.Printf("bit %d ", bit)
		// for i := 0; i < n; i++ {
		// 	fmt.Print(s[i])
		// }
		// fmt.Println()
	}
}
