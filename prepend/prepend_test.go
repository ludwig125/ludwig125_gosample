package main

import (
	"fmt"
	"testing"
)

var result []int

const size = 32

func prependSimple() {
	data := make([]int, 0, size)
	for i := 0; i < size; i++ {
		data = append([]int{i}, data...)
	}
	result = data
}

func prependWithCopy() {
	data := make([]int, 0, size)
	for i := 0; i < size; i++ {
		data = append(data, 0)
		copy(data[1:], data)
		data[0] = i
	}
	result = data
}

func TestPrependSimple(t *testing.T) {
	prependSimple()
	fmt.Println("prependSimple:  ", result)
}

func TestPrependWithCopy(t *testing.T) {
	prependWithCopy()
	fmt.Println("prependWithCopy:", result)
}
func BenchmarkPrependSimple(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prependSimple()
	}
}

func BenchmarkPrependWithCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prependWithCopy()
	}
}
