package main

import "fmt"

func main() {
	p := &[3]int{1, 2, 3}
	fmt.Println((*p)[0]) // 1
	fmt.Println((*p)[1]) // 2
	fmt.Println(p[2])    // 3
}
