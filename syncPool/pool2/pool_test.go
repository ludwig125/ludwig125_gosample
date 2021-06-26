package main

import (
	"sync"
	"testing"
)

// https://medium.com/swlh/go-the-idea-behind-sync-pool-32da5089df72

type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() interface{} { return new(Person) },
}

func BenchmarkWithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = new(Person)
			p.Age = 23
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = personPool.Get().(*Person)
			p.Age = 23
			personPool.Put(p)
		}
	}
}

var i int64

func BenchmarkPool(b *testing.B) {
	var p sync.Pool
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i = 100000000000000000
			p.Put(&i)
			p.Get()
		}
	})
}

var j int64

func BenchmarkAllocation(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			j = 100000000000000000
		}
	})
}

// [~/go/src/github.com/ludwig125/ludwig125_gosample/syncPool/pool2] $go test -bench . -benchmem
// goos: linux
// goarch: amd64
// pkg: github.com/ludwig125/ludwig125_gosample/syncPool/pool2
// BenchmarkWithoutPool-8              7254            145709 ns/op           80000 B/op      10000 allocs/op
// BenchmarkWithPool-8                 7470            150855 ns/op               0 B/op          0 allocs/op
// BenchmarkPool-8                 100000000               10.3 ns/op             0 B/op          0 allocs/op
// BenchmarkAllocation-8           502958653                2.44 ns/op            0 B/op          0 allocs/op
// PASS
// ok      github.com/ludwig125/ludwig125_gosample/syncPool/pool2  4.739s
