package main

import "fmt"

func Add(a, b int) int {
	defer fmt.Println("this is result")
	defer fmt.Println("this is result2")
	defer fmt.Println("this is result3")
	defer fmt.Println("this is result4")

	fmt.Println("Let's calicurate!")
	return a + b
}

func main() {
	x, y := 10, 30
	fmt.Println(Add(x, y))

}
