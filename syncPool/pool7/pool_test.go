package main

import (
	"bytes"
	"sort"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//https://yoru9zine.hatenablog.com/entry/2015/12/05/143414

func func1(in string) []byte {
	buf := &bytes.Buffer{}
	buf.WriteString(in)

	return buf.Bytes()
}

var globalBuf = &bytes.Buffer{}
var globalMutex = sync.Mutex{}

func func2(in string) []byte {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	globalBuf.Reset()

	globalBuf.WriteString(in)
	return globalBuf.Bytes()
}

var globalPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func func3(in string) []byte {
	buf := globalPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		globalPool.Put(buf)
	}()

	buf.WriteString(in)
	return buf.Bytes()
}

var out []byte

func makeStr(s string, n int) string {
	var ss string
	for i := 0; i < n; i++ {
		ss += s
	}
	return ss
}

var inputStr string

func init() {
	inputStr = makeStr("12345", 100000)
}

// func BenchmarkPoolFunc1(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		out = func1(inputStr)
// 	}
// }

// func BenchmarkPoolFunc2(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		out = func2(inputStr)
// 	}
// }

// func BenchmarkPoolFunc3(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		out = func3(inputStr)
// 	}
// }

// func BenchmarkPoolFunc1Conc(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		out = doFuncConcurrently(inputStr, 1000, func1)
// 	}
// }

// func BenchmarkPoolFunc2Conc(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		out = doFuncConcurrently(inputStr,1000,  func2)
// 	}
// }

// func BenchmarkPoolFunc3Conc(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		out = doFuncConcurrently(inputStr, 1000, func3)
// 	}
// }

func Benchmark(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	cases := map[string]struct {
		fn func(string) []byte
	}{
		"func1": {
			fn: func1,
		},
		"func2": {
			fn: func2,
		},
		"func3": {
			fn: func3,
		},
	}
	for name, tt := range cases {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				out = tt.fn(inputStr)
			}
		})
	}
}

func BenchmarkConc(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	cases := []struct {
		name string
		fn   func(string) []byte
		sep  int
	}{
		{
			name: "func1_sep10",
			fn:   func1,
			sep:  10,
		},
		{
			name: "func1_sep100000",
			fn:   func1,
			sep:  100000,
		},
		{
			name: "func1_sep1000000000",
			fn:   func1,
			sep:  1000000000,
		},
		{
			name: "func2_sep10",
			fn:   func2,
			sep:  10,
		},
		{
			name: "func2_sep100000",
			fn:   func2,
			sep:  100000,
		},
		{
			name: "func2_sep1000000000",
			fn:   func2,
			sep:  1000000000,
		},
		{
			name: "func3_sep10",
			fn:   func3,
			sep:  10,
		},
		{
			name: "func3_sep100000",
			fn:   func3,
			sep:  100000,
		},
		{
			name: "func3_sep1000000000",
			fn:   func3,
			sep:  1000000000,
		},
	}
	for _, tt := range cases {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				out = doFuncConcurrently(inputStr, tt.sep, tt.fn)
			}
		})
		b.Run(tt.name+"_withoutSorting", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				out = doFuncConcurrentlyWithoutSorting(inputStr, tt.sep, tt.fn)
			}
		})
	}
}

func doFuncConcurrentlyWithoutSorting(in string, sep int, fn func(string) []byte) []byte {
	l := separeteByN(sep, in)
	ch := make(chan []byte, len(l))
	for _, s := range l {
		go func(s string) {
			ch <- fn(s)
		}(s)
	}

	var ress []byte
	for i := 0; i < len(l); i++ {
		r := <-ch
		ress = append(ress, r...)
	}
	return ress
}

type numberedBytes struct {
	b []byte
	n int // lの何番目か
}

func doFuncConcurrently(in string, sep int, fn func(string) []byte) []byte {
	l := separeteByN(sep, in)
	ch := make(chan numberedBytes, len(l))
	for num, s := range l {
		go func(num int, s string) {
			ch <- numberedBytes{
				b: fn(s),
				n: num,
			}
		}(num, s)
	}

	var ress []numberedBytes
	for i := 0; i < len(l); i++ {
		r := <-ch
		ress = append(ress, r)
	}
	return sortAndMerge(ress)
}

// appendだと超遅い
// func separeteByN(n int, s string) []string {
// 	size := len(s) / n

// 	l := make([]string, 0, size)
// 	for {
// 		if len(s) >= n {
// 			l = append(l, s[:n])
// 			s = s[n:]
// 		} else {
// 			l = append(l, s)
// 			break
// 		}
// 	}
// 	return l
// }

func separeteByN(n int, s string) []string {
	size := len(s) / n

	l := make([]string, size)
	for i := 0; i < size; i++ {
		if len(s) >= n {
			l[i] = s[:n]
			s = s[n:]
		} else {
			l[i] = s[:n]
			break
		}
	}
	return l
}

func sortAndMerge(ress []numberedBytes) []byte {
	sort.Slice(ress, func(i, j int) bool { return ress[i].n < ress[j].n })
	var bs []byte
	for _, r := range ress {
		bs = append(bs, r.b...)
	}
	return bs
}

func TestFunc(t *testing.T) {
	in := inputStr
	want := []byte(inputStr)

	got := func1(in)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("got: %v,want: %v, diff: %s", got, want, diff)
	}
	got = doFuncConcurrently(in, 1000, func1)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("got: %v,want: %v, diff: %s", got, want, diff)
	}

	got = func2(in)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("got: %v,want: %v, diff: %s", got, want, diff)
	}
	got = doFuncConcurrently(in, 1000, func2)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("got: %v,want: %v, diff: %s", got, want, diff)
	}

	got = func3(in)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("got: %v,want: %v, diff: %s", got, want, diff)
	}
	got = doFuncConcurrently(in, 1000, func3)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("got: %v,want: %v, diff: %s", got, want, diff)
	}
}

// [~/go/src/github.com/ludwig125/ludwig125_gosample/syncPool/pool7] $go test -bench . -benchmem -benchtime=3s
// goos: linux
// goarch: amd64
// pkg: github.com/ludwig125/ludwig125_gosample/syncPool/pool7
// Benchmark/func1-8          25179            136615 ns/op          507906 B/op          1 allocs/op
// Benchmark/func2-8         165476             20506 ns/op               0 B/op          0 allocs/op
// Benchmark/func3-8         172185             19955 ns/op               2 B/op          0 allocs/op
// BenchmarkConc/func1_sep10-8                  108          33935922 ns/op        17801065 B/op      50218 allocs/op
// BenchmarkConc/func1_sep10_withoutSorting-8                   177          20949328 ns/op         8162477 B/op      50097 allocs/op
// BenchmarkConc/func1_sep100000-8                             5768            625967 ns/op         2147233 B/op         20 allocs/op
// BenchmarkConc/func1_sep100000_withoutSorting-8              5913            603187 ns/op         2146610 B/op         13 allocs/op
// BenchmarkConc/func1_sep1000000000-8                     36551768                99.8 ns/op            96 B/op          1 allocs/op
// BenchmarkConc/func1_sep1000000000_withoutSorting-8      48794941                74.4 ns/op            96 B/op          1 allocs/op
// BenchmarkConc/func2_sep10-8                                  100          31577774 ns/op        14595816 B/op        315 allocs/op
// BenchmarkConc/func2_sep10_withoutSorting-8                   139          25835493 ns/op         5190420 B/op       2664 allocs/op
// BenchmarkConc/func2_sep100000-8                             7720            463059 ns/op         1614753 B/op         15 allocs/op
// BenchmarkConc/func2_sep100000_withoutSorting-8              7780            461646 ns/op         1614129 B/op          8 allocs/op
// BenchmarkConc/func2_sep1000000000-8                     30655568                99.9 ns/op            96 B/op          1 allocs/op
// BenchmarkConc/func2_sep1000000000_withoutSorting-8      41327768                73.9 ns/op            96 B/op          1 allocs/op
// BenchmarkConc/func3_sep10-8                                  100          31632930 ns/op        14572289 B/op         66 allocs/op
// BenchmarkConc/func3_sep10_withoutSorting-8                   187          19201911 ns/op         4938086 B/op         34 allocs/op
// BenchmarkConc/func3_sep100000-8                             7722            466696 ns/op         1627515 B/op         15 allocs/op
// BenchmarkConc/func3_sep100000_withoutSorting-8              7827            460435 ns/op         1626218 B/op          8 allocs/op
// BenchmarkConc/func3_sep1000000000-8                     30274360               101 ns/op              96 B/op          1 allocs/op
// BenchmarkConc/func3_sep1000000000_withoutSorting-8      42524690                73.7 ns/op            96 B/op          1 allocs/op
// PASS
// ok      github.com/ludwig125/ludwig125_gosample/syncPool/pool7  113.795s
// [~/go/src/github.com/ludwig125/ludwig125_gosample/syncPool/pool7] $
