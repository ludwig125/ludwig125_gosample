package main

import "fmt"

type Vertex struct {
	a, b int
}

func main() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40, -74,
	}
	fmt.Println(m["Bell Labs"])
}
