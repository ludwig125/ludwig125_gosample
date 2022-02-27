package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
	"unsafe"

	"github.com/google/go-cmp/cmp"
)

const (
	// CtrlA ^A
	CtrlA string = "\x01"
	// CtrlB ^B
	CtrlB string = "\x02"
	// CtrlC ^C
	CtrlC string = "\x03"

	IoidLength = 10
)

// ConditionListID is ...
type ConditionListID string

// SimConditionListID is ...
type SimConditionListID string

// ConditionListInfo contains
type ConditionListInfo struct {
	IOs        []string
	Expire     time.Duration
	DeleteFlag bool
	ActiveFlag bool
}

// SimConditionListInfo contains
type SimConditionListInfo struct {
	TLID ConditionListID
	ConditionListInfo
	Threshold float64
}

// Get is nealy original function.
// https://ghe.corp.yahoo.co.jp/anemos/user-service/blob/25e84653bd62547d8f6298d5617f8fc82e6519e2/gateways/tsv/target_list_info_dictionary.go#L165
func Get(b []byte) (ConditionListInfo, error) {
	// b, found := dict.d.Get(string(id))
	// if !found {
	// 	return rtg.ConditionListInfo{}, nil
	// }
	s := *(*string)(unsafe.Pointer(&b))

	data := strings.Split(s, CtrlA)

	if len(data) != 4 {
		return ConditionListInfo{}, fmt.Errorf("failed to split")
	}

	expire, err := strconv.ParseInt(data[0], 10, 64)
	if err != nil {
		return ConditionListInfo{}, err
	}

	deleteFlag := data[1] == "1"

	activeFlag := true
	if data[2] == "0" {
		activeFlag = false
	}

	ioids := strings.Split(data[3], CtrlB)

	info := ConditionListInfo{
		IOs:        ioids,
		Expire:     time.Duration(expire),
		DeleteFlag: deleteFlag,
		ActiveFlag: activeFlag,
	}

	return info, nil
}

func GetMod(b []byte) (ConditionListInfo, error) {
	s := *(*string)(unsafe.Pointer(&b))

	data := strings.Split(s, CtrlA)

	if len(data) != 4 {
		return ConditionListInfo{}, fmt.Errorf("failed to split")
	}

	expire, err := strconv.ParseInt(data[0], 10, 64)
	if err != nil {
		return ConditionListInfo{}, err
	}

	deleteFlag := data[1] == "1"

	// activeFlag := true
	// if data[2] == "0" {
	// 	activeFlag = false
	// }
	activeFlag := data[2] != "0"

	// ioids := strings.Split(data[3], CtrlB)
	ioidsStr := data[3]
	length := len(ioidsStr)

	ioids := make([]string, 0, length/IoidLength)
	for i := IoidLength; i <= length; i += IoidLength {
		ioids = append(ioids, ioidsStr[:IoidLength])
		ioidsStr = ioidsStr[IoidLength:]
	}
	info := ConditionListInfo{
		IOs:        ioids,
		Expire:     time.Duration(expire),
		DeleteFlag: deleteFlag,
		ActiveFlag: activeFlag,
	}

	return info, nil
}

var conditionListIoidsPool = sync.Pool{
	New: func() interface{} {
		s := make([]string, 0, 0)
		return &s
	},
}

func GetModSyncPool(b []byte) (ConditionListInfo, error) {
	s := *(*string)(unsafe.Pointer(&b))

	data := strings.Split(s, CtrlA)

	if len(data) != 4 {
		return ConditionListInfo{}, fmt.Errorf("failed to split")
	}

	expire, err := strconv.ParseInt(data[0], 10, 64)
	if err != nil {
		return ConditionListInfo{}, err
	}

	deleteFlag := data[1] == "1"

	activeFlag := data[2] != "0"

	ioidsStr := data[3]
	length := len(ioidsStr)

	// ioids := make([]string, 0, length/IoidLength)

	// sync.Poolを使った方法
	ioidsPtr := conditionListIoidsPool.Get().(*[]string)
	// ioids := *ioidsPtr
	// ioids = ioids[:0] // sliceの容量はそのまま値だけクリアする

	// ioids := make([]string, len(*ioidsPtr))
	// copy(ioids, *ioidsPtr)

	// ioids := append((*ioidsPtr)[:0:0], (*ioidsPtr)...)

	ioids := make([]string, 0, len(*ioidsPtr))
	copy(ioids, (*ioidsPtr)[:0])

	for i := IoidLength; i <= length; i += IoidLength {
		// ioids = append(ioids, ioidsStr[:IoidLength])
		ioid := ioidsStr[:IoidLength]
		ioids = append(ioids, ioid)
		ioidsStr = ioidsStr[IoidLength:]
	}
	info := ConditionListInfo{
		IOs:        ioids,
		Expire:     time.Duration(expire),
		DeleteFlag: deleteFlag,
		ActiveFlag: activeFlag,
	}

	*ioidsPtr = ioids
	conditionListIoidsPool.Put(ioidsPtr)

	// ioidsPtr := conditionListIoidsPool.Get().(*[]string)
	// ioids := *ioidsPtr

	// ioids = ioids[:0] // sliceの容量はそのまま値だけクリアする
	// for i := IoidLength; i <= length; i += IoidLength {
	// 	ioids = append(ioids, ioidsStr[:IoidLength])
	// 	ioidsStr = ioidsStr[IoidLength:]
	// }
	// info := ConditionListInfo{
	// 	IOs:        ioids,
	// 	Expire:     time.Duration(expire),
	// 	DeleteFlag: deleteFlag,
	// 	ActiveFlag: activeFlag,
	// }

	// *ioidsPtr = ioids
	// conditionListIoidsPool.Put(ioidsPtr)
	// 上の書き方だとこのBenchmark
	// BenchmarkGetModSyncPool
	// BenchmarkGetModSyncPool-8        3086025               375.8 ns/op            64 B/op          1 allocs/op
	// BenchmarkGetModSyncPool-8        3162134               371.7 ns/op            64 B/op          1 allocs/op
	// BenchmarkGetModSyncPool-8        3237186               371.9 ns/op            64 B/op          1 allocs/op
	// BenchmarkGetModSyncPool-8        3235969               372.9 ns/op            64 B/op          1 allocs/op

	// 以下よりも上のほうが若干速かった
	// ioids := conditionListIoidsPool.Get().(*[]string)
	// (*ioids) = (*ioids)[:0] // sliceの容量はそのまま値だけクリアする
	// for i := IoidLength; i <= length; i += IoidLength {
	// 	(*ioids) = append((*ioids), ioidsStr[:IoidLength])
	// 	ioidsStr = ioidsStr[IoidLength:]
	// }
	// info := ConditionListInfo{
	// 	IOs:        *ioids,
	// 	Expire:     time.Duration(expire),
	// 	DeleteFlag: deleteFlag,
	// 	ActiveFlag: activeFlag,
	// }
	// conditionListIoidsPool.Put(ioids)
	// 上の書き方だとこのBenchmark
	// BenchmarkGetModSyncPool
	// BenchmarkGetModSyncPool-8        2891492               414.3 ns/op            64 B/op          1 allocs/op
	// BenchmarkGetModSyncPool-8        2872594               415.8 ns/op            64 B/op          1 allocs/op
	// BenchmarkGetModSyncPool-8        2831012               424.9 ns/op            64 B/op          1 allocs/op
	// BenchmarkGetModSyncPool-8        2803974               435.6 ns/op            64 B/op          1 allocs/op

	return info, nil
}

func GetMod2(b []byte) (ConditionListInfo, error) {
	s := *(*string)(unsafe.Pointer(&b))

	data := strings.Split(s, CtrlA)

	if len(data) != 4 {
		return ConditionListInfo{}, fmt.Errorf("failed to split")
	}

	expire, err := strconv.ParseInt(data[0], 10, 64)
	if err != nil {
		return ConditionListInfo{}, err
	}

	deleteFlag := data[1] == "1"

	// activeFlag := true
	// if data[2] == "0" {
	// 	activeFlag = false
	// }
	activeFlag := data[2] != "0"

	// ioids := strings.Split(data[3], CtrlB)
	ioidsStr := data[3]
	length := len(ioidsStr)

	ioids := make([]string, 0, length/IoidLength)
	for i := IoidLength; i <= length; i += IoidLength {
		ioids = append(ioids, ioidsStr[:IoidLength])
		ioidsStr = ioidsStr[IoidLength:]
	}
	info := ConditionListInfo{
		IOs:        ioids,
		Expire:     time.Duration(expire),
		DeleteFlag: deleteFlag,
		ActiveFlag: activeFlag,
	}

	return info, nil
}

// target_list_info_dictionary.goのLoadで作成するデータを再現したもの
// https://ghe.corp.yahoo.co.jp/anemos/user-service/blob/25e84653bd62547d8f6298d5617f8fc82e6519e2/gateways/tsv/target_list_info_dictionary.go#L121-L128
func dummyConditionList(n int) []byte {
	var ioid string
	for i := 0; i < n; i++ {
		tmp := fmt.Sprintf("12%08d", i)
		if i != 0 {
			tmp = CtrlB + tmp
		}
		ioid += tmp
	}
	span := int64(12345)
	day := int64(1)
	deleteFlag := "0"
	activeFlag := "1"

	res := strconv.FormatInt(span*int64(day), 10) + CtrlA + deleteFlag + CtrlA + activeFlag + CtrlA + ioid
	return []byte(res)
}

// dummyConditionListのioidのCtrlB区切りをやめたもの
func dummyConditionListMod(n int) []byte {
	var ioids string
	for i := 0; i < n; i++ {
		ioids += fmt.Sprintf("12%08d", i)
	}
	span := int64(12345)
	day := int64(1)
	deleteFlag := "0"
	activeFlag := "1"

	res := strconv.FormatInt(span*int64(day), 10) + CtrlA + deleteFlag + CtrlA + activeFlag + CtrlA + ioids
	return []byte(res)
}

func dummyConditionListMod2(n int) []byte {
	var ioids string
	for i := 0; i < n; i++ {
		ioids += fmt.Sprintf("12%08d", i)
	}
	span := int64(12345)
	day := int64(1)
	deleteFlag := "0"
	activeFlag := "1"

	res := strconv.FormatInt(span*int64(day), 10) + CtrlA + deleteFlag + CtrlA + activeFlag + CtrlA + ioids

	return []byte(res)
}

func dummyConditionListMod3(n int) []byte {
	var ioids string
	for i := 0; i < n; i++ {
		ioids += fmt.Sprintf("12%08d", i)
	}
	span := int64(12345)
	day := int64(1)
	deleteFlag := "0"
	activeFlag := "1"

	res := strconv.FormatInt(span*int64(day), 10) + CtrlA + deleteFlag + CtrlA + activeFlag + CtrlA + ioids

	return []byte(res)
}

var DummyConditionList = dummyConditionList(100)
var DummyConditionListMod = dummyConditionListMod(100)
var DummyConditionListMod2 = dummyConditionListMod2(100)

func TestConditionListInfoDictionaryGet(t *testing.T) {
	var wg sync.WaitGroup
	num := 2
	wg.Add(num)
	for i := 0; i < num; i++ {
		i := i
		go func() {
			defer wg.Done()

			dummy := dummyConditionList(50 + i)
			// infos, err := Get(dummy)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			dummyMod := dummyConditionListMod(50 + i)
			dummyMod2 := dummyConditionListMod2(50 + i)

			var err error
			var infos ConditionListInfo
			fmt.Println("AllocsPerRun: ", int(testing.AllocsPerRun(100, func() {
				infos, err = Get(dummy)
				if err != nil {
					log.Fatal(err)
				}
			})))

			var infos2 ConditionListInfo
			fmt.Println("AllocsPerRun2:", int(testing.AllocsPerRun(100, func() {
				infos2, err = GetMod(dummyMod)
				if err != nil {
					log.Fatal(err)
				}
			})))
			if diff := cmp.Diff(infos2, infos); diff != "" {
				t.Errorf("(-infos2 +infos)\n%s", diff)
			}

			var infos3 ConditionListInfo
			fmt.Println("AllocsPerRun3:", int(testing.AllocsPerRun(100, func() {
				infos3, err = GetModSyncPool(dummyMod)
				if err != nil {
					log.Fatal(err)
				}
			})))
			if diff := cmp.Diff(infos3, infos); diff != "" {
				t.Errorf("(-infos3 +infos)\n%s", diff)
			}

			var infos4 ConditionListInfo
			fmt.Println("AllocsPerRun4:", int(testing.AllocsPerRun(100, func() {
				infos4, err = GetMod2(dummyMod2)
				if err != nil {
					log.Fatal(err)
				}
			})))
			if diff := cmp.Diff(infos4, infos); diff != "" {
				t.Errorf("(-infos4 +infos)\n%s", diff)
			}
		}()
	}
	wg.Wait()

	heapfile, err := os.Create("heap.pprof")
	if err != nil {
		panic(err)
	}
	err = pprof.WriteHeapProfile(heapfile)
	if err != nil {
		panic(err)
	}
	defer heapfile.Close()
}

var Result ConditionListInfo

func BenchmarkGet(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var res ConditionListInfo
	for n := 0; n < b.N; n++ {
		res, _ = Get(DummyConditionList)
	}
	Result = res
}

func BenchmarkGetMod(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var res ConditionListInfo
	for n := 0; n < b.N; n++ {
		res, _ = GetMod(DummyConditionListMod)
	}
	Result = res
}

func BenchmarkGetModSyncPool(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var res ConditionListInfo
	for n := 0; n < b.N; n++ {
		res, _ = GetModSyncPool(DummyConditionListMod)
	}
	Result = res
}

func BenchmarkGetMod2(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var res ConditionListInfo
	for n := 0; n < b.N; n++ {
		res, _ = GetMod2(DummyConditionListMod2)
	}
	Result = res
}
