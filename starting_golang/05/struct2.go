package main

import "fmt"

type Point struct{ X, Y int }

// 値渡し
func swap(p Point) {
	tmp := p.X
	p.X = p.Y
	p.Y = tmp
	fmt.Printf("swap %d %d\n", p.X, p.Y)
}

// ポインタ渡し
func swap_pointer(p *Point) {
	tmp := p.X
	p.X = p.Y
	p.Y = tmp
	fmt.Printf("swap %d %d\n", p.X, p.Y)
}

func main() {
	p := Point{X: 10, Y: 20}
	fmt.Printf("before %d %d\n", p.X, p.Y) // 10 20
	swap(p)                                // 20 10
	fmt.Printf("after %d %d\n", p.X, p.Y)  // 10 20 変わってない

	swap_pointer(&p)                      // 20 10
	fmt.Printf("after %d %d\n", p.X, p.Y) // 20 10 変わっている
}
