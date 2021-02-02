package main

import "fmt"

func inc(p *int) {
	*p++
}

func main() {
	var i int
	p := &i
	fmt.Printf("%T\n", p) // *int
	pp := &p
	fmt.Printf("%T\n", pp) // **int

	i = 5
	fmt.Println(i) // 5
	*p = 10
	fmt.Println(i) // 10
	inc(p)
	fmt.Println(i) // 11
	inc(&i)
	fmt.Println(i) // 12
}
