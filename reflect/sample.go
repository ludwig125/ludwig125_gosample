package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X int
	Y int
}

func main() {
	p := Point{X: 10, Y: 5}
	fmt.Printf("reflect.ValueOf(p): (p is value) %T\n", reflect.ValueOf(p))

	p2 := &Point{X: 10, Y: 5}
	fmt.Printf("reflect.ValueOf(p2): (p2 is pointer) %T\n", reflect.ValueOf(p2))

	p3 := &Point{X: 10, Y: 5}
	fmt.Printf("reflect.ValueOf(p3).Elem(): (p3 is pointer) %T\n", reflect.ValueOf(p3).Elem())
}
