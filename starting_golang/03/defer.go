package main

import "fmt"

func main() {
	// defer は逆順に実行される
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("A")
}
