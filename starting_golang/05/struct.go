package main

import "fmt"

type T struct {
	int
	float64
	string
	N int
}

func main() {
	t := T{1, 3.14, "文字列", 10}
	fmt.Println(t.int)
	fmt.Println(t.float64)
	fmt.Println(t.string)
	fmt.Println(t.N)
}
