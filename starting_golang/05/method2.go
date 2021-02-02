package main

import "fmt"

type MyInt int

func (m MyInt) Plus(i int) int {
	return int(m) + i
}

func main() {
	fmt.Println(MyInt(2).Plus(3))
}
