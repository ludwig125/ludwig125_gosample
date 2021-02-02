package main

import (
	"fmt"
	"reflect"
	"testing"
)

func convertStringSlicesToInterfaceSlices1(sss [][]string) [][]interface{} {
	var iss [][]interface{}
	for _, ss := range sss {
		var is []interface{}
		for _, s := range ss {
			is = append(is, s)
		}
		iss = append(iss, is)
	}
	return iss
}

func convertStringSlicesToInterfaceSlices2(sss [][]string) [][]interface{} {
	iss := make([][]interface{}, 0, len(sss))
	for _, ss := range sss {
		is := make([]interface{}, 0, len(ss))
		for _, s := range ss {
			is = append(is, s)
		}
		iss = append(iss, is)
	}
	return iss
}

func convertStringSlicesToInterfaceSlices3(sss [][]string) [][]interface{} {
	iss := make([][]interface{}, len(sss))
	for i, ss := range sss {
		is := make([]interface{}, len(ss))
		for j, s := range ss {
			is[j] = s
		}
		iss[i] = is
	}
	return iss
}

func convertStringSlicesToInterfaceSlices4(sss [][]string) [][]interface{} {
	iss := make([][]interface{}, len(sss))
	for i, ss := range sss {
		iss[i] = interfaceSlice(ss)
	}
	return iss
}

func interfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("interfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

func makeStringSlices(n, m int) [][]string {
	var ress [][]string
	for i := 0; i < n; i++ {
		var res []string
		for j := 0; j < m; j++ {
			res = append(res, fmt.Sprintf("%d", j))
		}
		ress = append(ress, res)
	}
	return ress
}

func makeInterfaceSlices(n, m int) [][]interface{} {
	var ress [][]interface{}
	for i := 0; i < n; i++ {
		var res []interface{}
		for j := 0; j < m; j++ {
			res = append(res, fmt.Sprintf("%d", j))
		}
		ress = append(ress, res)
	}
	return ress
}

func TestConvertStringSlicesToInterfaceSlices(t *testing.T) {
	inputSss := makeStringSlices(3, 3)
	wantIss := makeInterfaceSlices(3, 3)
	t.Log("inputSss:", inputSss)
	t.Log("wantIss:", wantIss)
	if !reflect.DeepEqual(convertStringSlicesToInterfaceSlices1(inputSss), wantIss) {
		t.Error("failed to convertStringSlicesToInterfaceSlices1")
	}
	if !reflect.DeepEqual(convertStringSlicesToInterfaceSlices2(inputSss), wantIss) {
		t.Error("failed to convertStringSlicesToInterfaceSlices2")
	}
	if !reflect.DeepEqual(convertStringSlicesToInterfaceSlices3(inputSss), wantIss) {
		t.Error("failed to convertStringSlicesToInterfaceSlices3")
	}
	if !reflect.DeepEqual(convertStringSlicesToInterfaceSlices4(inputSss), wantIss) {
		t.Error("failed to convertStringSlicesToInterfaceSlices4")
	}
}

var N = 10000
var result [][]interface{}

func BenchmarkTestConvertSlices1(b *testing.B) {
	inputSss := makeStringSlices(10, 10)
	b.ResetTimer()
	var res [][]interface{}
	for i := 0; i < N; i++ {
		res = convertStringSlicesToInterfaceSlices1(inputSss)
	}
	result = res
}

func BenchmarkTestConvertSlices2(b *testing.B) {
	inputSss := makeStringSlices(10, 10)
	b.ResetTimer()
	var res [][]interface{}
	for i := 0; i < N; i++ {
		res = convertStringSlicesToInterfaceSlices2(inputSss)
	}
	result = res
}

func BenchmarkTestConvertSlices3(b *testing.B) {
	inputSss := makeStringSlices(10, 10)
	b.ResetTimer()
	var res [][]interface{}
	for i := 0; i < N; i++ {
		res = convertStringSlicesToInterfaceSlices3(inputSss)
	}
	result = res
}

func BenchmarkTestConvertSlices4(b *testing.B) {
	inputSss := makeStringSlices(10, 10)
	b.ResetTimer()
	var res [][]interface{}
	for i := 0; i < N; i++ {
		res = convertStringSlicesToInterfaceSlices4(inputSss)
	}
	result = res
}
