package main

import "fmt"

type Point struct{ X, Y int }

func (p *Point) Render() {
	fmt.Printf("%d %d\n", p.X, p.Y)
}

// 値渡し
func (p Point) swap() {
	tmp := p.X
	p.X = p.Y
	p.Y = tmp
	fmt.Printf("swap atai %d %d\n", p.X, p.Y)
}

// ポインタ渡し
func (p *Point) swapp() {
	tmp := p.X
	p.X = p.Y
	p.Y = tmp
	fmt.Printf("swap sansyo %d %d\n", p.X, p.Y)
}

func main() {
	p := &Point{X: 5, Y: 12}
	p.Render()

	p2 := Point{X: 10, Y: 20}
	fmt.Printf("before %d %d\n", p2.X, p2.Y)
	p2.swap()
	fmt.Printf("after %d %d\n", p2.X, p2.Y)
	p2.swapp()
	fmt.Printf("after %d %d\n", p2.X, p2.Y)
}
