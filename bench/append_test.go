package main

import (
	"fmt"
	"testing"
)

var N = 1000000

func BenchmarkTest1(b *testing.B) {
	b.ResetTimer()
	list := []string{}
	for i := 0; i < N; i++ {
		list = append(list, fmt.Sprintf("%d", i))
	}
}

func BenchmarkTest2(b *testing.B) {
	b.ResetTimer()
	list := make([]string, N)
	for i := 0; i < N; i++ {
		list = append(list, fmt.Sprintf("%d", i))
	}
}

func BenchmarkTest3(b *testing.B) {
	b.ResetTimer()
	list := make([]string, 0, N)
	for i := 0; i < N; i++ {
		list = append(list, fmt.Sprintf("%d", i))
	}
}

func BenchmarkTest4(b *testing.B) {
	b.ResetTimer()
	list := make([]string, N)
	for i := 0; i < N; i++ {
		list[i] = fmt.Sprintf("%d", i)
	}
}
