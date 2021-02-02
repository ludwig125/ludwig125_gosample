package concat

import "testing"

func BenchmarkConcat(b *testing.B) {
	b.ResetTimer()

	var ss []string
	for n := 0; n < 100; n++ {
		ss = append(ss, "foo")
	}

	for i := 0; i < b.N; i++ {
		concat(ss...)
	}
}

func BenchmarkConcatByteBuffer(b *testing.B) {
	//b.ResetTimer()

	var ss []string
	for n := 0; n < 100; n++ {
		ss = append(ss, "foo")
	}

	for i := 0; i < b.N; i++ {
		concatByteBuffer(ss...)
	}
}
