package main

import "fmt"

//任意の個数のint型の合計値を返す
func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

// 配列型は値渡し
func pow(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

// 参照型（スライス）は参照渡し
func pow2(s []int) {
	for i, v := range s {
		s[i] = v * v
	}
	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	// スライスを直接引数に入れることもできる
	s := []int{1, 2, 3}
	fmt.Println(sum(s...))

	// 配列型
	a := [3]int{1, 2, 3}
	pow(a)
	fmt.Println(a) // [1, 2, 3] 変わらない

	pow2(s)
	fmt.Println(s) // [1, 4, 9] 参照側を渡しているので自乗されている

}
