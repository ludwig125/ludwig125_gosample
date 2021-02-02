package main

import "fmt"

func plus(x, y int) int {
	return x + y
}

func hello() {
	fmt.Println("Hello World")
	return
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func main() {
	fmt.Println(plus(1, 3))
	hello()
	q, r := div(10, 3)
	fmt.Printf("%d %d\n", q, r)

	// 無名関数
	f := func(x, y int) int { return x + y }
	fmt.Printf("%d\n", f(10, 5))

}
