package main

import "fmt"

func main() {
	s := make([]int, 10) // slice
	fmt.Println(s)
	s[0] = 1
	fmt.Println(s)
	s[1] = 2
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	arr := [3]int{1, 2, 3} // array
	fmt.Println(arr)
	fmt.Println(len(arr))

	s = append(s, 101)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s = append(s, 102, 103, 104)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	str := []string{"Apple", "Banana", "Cherry"}
	for i, v := range str {
		fmt.Printf("[%d] -> %s\n", i, v)
	}
}
