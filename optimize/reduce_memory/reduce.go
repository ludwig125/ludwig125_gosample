package main

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"math/rand"
	"os"
	"runtime/pprof"
	"strconv"
	"sync"
	"testing"
	"unsafe"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		// length of a sha256 hash
		b := make([]byte, 256)
		return &b
	},
}

var hashPool = sync.Pool{
	New: func() interface{} {
		return sha256.New()
	},
}

func foo(n int) string {
	// origin
	// var buf bytes.Buffer
	// x := strconv.Itoa(n)
	// for i := 0; i < 100000; i++ {
	// 	buf.WriteString(x)
	// }
	// sum := sha256.Sum256(buf.Bytes())

	// before sync.Pool
	// x := strconv.Itoa(n)
	// buf := make([]byte, 0, 100000*len(x))
	// for i := 0; i < 100000; i++ {
	// 	buf = append(buf, x...)
	// }
	// sum := sha256.Sum256(buf)

	// get buffer from pool
	bufptr := bufPool.Get().(*[]byte)
	defer bufPool.Put(bufptr)
	buf := *bufptr
	// reset buf
	buf = buf[:0]

	// get hash object from pool
	h := hashPool.Get().(hash.Hash)
	defer hashPool.Put(h)
	h.Reset()

	x := strconv.AppendInt(buf, int64(n), 10)
	for i := 0; i < 100000; i++ {
		h.Write(x)
	}
	// reset whatever strconv.AppendInt put in the buf
	buf = buf[:0]
	sum := h.Sum(buf)

	// var b []byte
	// b := make([]byte, 0, int(sum[0]))
	b := make([]byte, 0, 256)
	for i := 0; i < int(sum[0]); i++ {
		x := sum[(i*7+1)%len(sum)] ^ sum[(i*5+3)%len(sum)]
		// c := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 10)[x]
		// c := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"[x]
		c := "abcdefghijklmnopqrstuvwxyz"[x%26]
		b = append(b, c)
	}
	// return string(b)
	// return *(*string)(unsafe.Pointer(&b))
	sum = sum[:0] // reset sum
	sum = append(sum, b...)
	return *(*string)(unsafe.Pointer(&sum))
}

func main() {
	cpufile, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	err = pprof.StartCPUProfile(cpufile)
	if err != nil {
		panic(err)
	}
	defer cpufile.Close()
	defer pprof.StopCPUProfile()

	// ensure function output is accurate
	if foo(12345) == "aajmtxaattdzsxnukawxwhmfotnm" {
		fmt.Println("Test PASS")
	} else {
		fmt.Println("Test FAIL")
	}

	fmt.Println("Allocs:", int(testing.AllocsPerRun(100, func() {
		foo(rand.Int())
	})))
}
