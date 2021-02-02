package swrap

import (
	"testing"
)

func Fixture() []byte {
	return []byte{
		0x0, 0x1, 0x2, 0x3, 0x4,
		0x5, 0x6, 0x7, 0x8, 0x9,
	}
}
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := New(Fixture())
		sw.Add(0xFF)
	}
}

func BenchmarkLen(b *testing.B) {
	sw := New(Fixture())
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sw.Len()
	}
}
