package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		a, b := fmt.Sprintf("%d", i-1), fmt.Sprintf("%d", i)
		fmt.Printf("[%s %s] [% x % x] %t\n", a, b, a, b, a < b)
	}

	a, b := "90", "9!"
	fmt.Printf("[%s %s] [% x % x] %t\n", a, b, a, b, a < b)
}
