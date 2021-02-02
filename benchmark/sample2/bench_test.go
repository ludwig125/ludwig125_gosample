package main

// https://qiita.com/syossan27/items/148e33dd9da4ee3dc89b

import "testing"

func BenchmarkMemAllocBeforeCustom(b *testing.B) {
	n := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < n; j++ {
			s = s + "alice"
		}
	}
}

func BenchmarkMemAllocCustom(b *testing.B) {
	n := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]byte, 0)
		for j := 0; j < n; j++ {
			s = append(s, "alice"...)
		}
	}
}
