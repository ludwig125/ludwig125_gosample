package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		if i == 9 {
			fmt.Println()
		}
	} // 0 1 2 3 4 5 6 7 8 9

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------------------")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		i++
	}
}
