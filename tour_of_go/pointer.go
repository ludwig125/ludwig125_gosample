package main

import "fmt"

func main() {
	x, y := 10, 30

	p := &x
	fmt.Println(x)
	*p += 1
	fmt.Println(x)

	p = &y
	fmt.Println(y)
	*p += 1
	fmt.Println(y)
}
