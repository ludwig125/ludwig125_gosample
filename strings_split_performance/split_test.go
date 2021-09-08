package main

import (
	"fmt"
	"strings"
	"testing"
)

const Delim = "-"

func FirstIsLongest(n int) string {
	var ids string
	for i := 0; i < n; i++ {
		ids += fmt.Sprintf("1%05d", i)
	}
	return ids + Delim + "aaaaaa" + Delim + "bbbbbbbb" + Delim + "cccccccc"
}

func MiddleIsLongest(n int) string {
	var ids string
	for i := 0; i < n; i++ {
		ids += fmt.Sprintf("1%05d", i)
	}
	return "aaaaaa" + Delim + ids + Delim + "bbbbbbbb" + Delim + "cccccccc"
}

func LastIsLongest(n int) string {
	var ids string
	for i := 0; i < n; i++ {
		ids += fmt.Sprintf("1%05d", i)
	}
	return "aaaaaa" + Delim + "bbbbbbbb" + Delim + "cccccccc" + Delim + ids
}

func SplitStr(s string) []string {
	return strings.Split(s, Delim)
}

var first = FirstIsLongest(1000)
var middle = MiddleIsLongest(1000)
var last = LastIsLongest(1000)
var Result []string

func BenchmarkSplitStrFirst(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var r []string
	for n := 0; n < b.N; n++ {
		r = SplitStr(first)
	}
	Result = r
}

func BenchmarkSplitStrMiddle(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var r []string
	for n := 0; n < b.N; n++ {
		r = SplitStr(middle)
	}
	Result = r
}

func BenchmarkSplitStrLast(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var r []string
	for n := 0; n < b.N; n++ {
		r = SplitStr(last)
	}
	Result = r
}

// [~/go/src/github.com/ludwig125/ludwig125_gosample/strings_split_performance] $go test -bench . -count=4
// goos: linux
// goarch: amd64
// pkg: github.com/ludwig125/ludwig125_gosample/strings_split_performance
// cpu: Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz
// BenchmarkSplitStrFirst-8         3257382               354.2 ns/op            64 B/op          1 allocs/op
// BenchmarkSplitStrFirst-8         3345314               359.0 ns/op            64 B/op          1 allocs/op
// BenchmarkSplitStrFirst-8         3389498               347.8 ns/op            64 B/op          1 allocs/op
// BenchmarkSplitStrFirst-8         3364164               347.4 ns/op            64 B/op          1 allocs/op
// BenchmarkSplitStrMiddle-8        3229088               364.2 ns/op            64 B/op          1 allocs/op
// BenchmarkSplitStrMiddle-8        3156192               366.3 ns/op            64 B/op          1 allocs/op
// BenchmarkSplitStrMiddle-8        3217290               372.4 ns/op            64 B/op          1 allocs/op
// BenchmarkSplitStrMiddle-8        3176865               367.2 ns/op            64 B/op          1 allocs/op
// BenchmarkSplitStrLast-8          5102700               226.2 ns/op            64 B/op          1 allocs/op
// BenchmarkSplitStrLast-8          5059629               227.0 ns/op            64 B/op          1 allocs/op
// BenchmarkSplitStrLast-8          5017147               225.6 ns/op            64 B/op          1 allocs/op
// BenchmarkSplitStrLast-8          5072168               226.6 ns/op            64 B/op          1 allocs/op
// PASS
// ok      github.com/ludwig125/ludwig125_gosample/strings_split_performance       17.979s
