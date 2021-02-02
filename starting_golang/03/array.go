package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", a)
	fmt.Println(a[0])

	b := [...]int{1, 2, 3}
	fmt.Printf("%v\n", b)
	b[1] = 10
	fmt.Printf("%v\n", b)
}
