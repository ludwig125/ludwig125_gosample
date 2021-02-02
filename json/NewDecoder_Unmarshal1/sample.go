package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// JSON EncodeしてDecodeすると中身が微妙に変わる

// Sample is struct for json parse.
type Sample struct {
	Price   float64 `json:",string"`
	Ratio15 float64 `json:",string"`
	Ratio16 float64 `json:",string"`
}

func main() {
	s := `{"price":"1.0", "ratio15":"1.000000000000001", "ratio16":"1.0000000000000001"}`
	fmt.Printf("before %+v\n\n", s)
	fn1(s)
	fn2(s)
	fn3(s)
	fn4(s)
}
func fn1(s string) {
	var x Sample
	err := json.Unmarshal([]byte(s), &x)
	if err != nil {
		log.Fatalf("%+v err: %v\n", x, err)
	}
	fmt.Printf("after Unmarshal %+v\n", x)

	buf := &bytes.Buffer{}
	// Encodeを用いたJson変換
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	encoder.Encode(x)
	fmt.Printf("after Encode %+v\n", string(buf.Bytes()))
}

func fn2(s string) {
	d := json.NewDecoder(strings.NewReader(s))
	d.UseNumber()
	var x Sample
	if err := d.Decode(&x); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("after decode %+v\n", x)

	buf := &bytes.Buffer{}
	// Encodeを用いたJson変換
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	encoder.Encode(x)
	fmt.Printf("after Encode %+v\n", string(buf.Bytes()))
}
func fn3(s string) {
	var x Sample
	err := json.Unmarshal([]byte(s), &x)
	if err != nil {
		log.Fatalf("%+v err: %v\n", x, err)
	}
	fmt.Printf("after Unmarshal %+v\n", x)

	buf := &bytes.Buffer{}
	// Encodeを用いたJson変換
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	encoder.Encode(x)
	fmt.Printf("after Encode %+v\n", string(buf.Bytes()))
}

func fn4(s string) {
	d := json.NewDecoder(strings.NewReader(s))
	d.UseNumber()
	var x interface{}
	if err := d.Decode(&x); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("after decode %+v\n", x)

	buf := &bytes.Buffer{}
	// Encodeを用いたJson変換
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	encoder.Encode(x)
	fmt.Printf("after Encode %+v\n", string(buf.Bytes()))
}

// 実行結果
// interface型で展開すると元の小数以下の数字も保持されている

// before {"price":"1.0", "ratio15":"1.000000000000001", "ratio16":"1.0000000000000001"}

// after Unmarshal {Price:1 Ratio15:1.000000000000001 Ratio16:1}
// after Encode {"Price":"1","Ratio15":"1.000000000000001","Ratio16":"1"}

// after decode {Price:1 Ratio15:1.000000000000001 Ratio16:1}
// after Encode {"Price":"1","Ratio15":"1.000000000000001","Ratio16":"1"}

// after Unmarshal {Price:1 Ratio15:1.000000000000001 Ratio16:1}
// after Encode {"Price":"1","Ratio15":"1.000000000000001","Ratio16":"1"}

// after decode map[price:1.0 ratio15:1.000000000000001 ratio16:1.0000000000000001]
// after Encode {"price":"1.0","ratio15":"1.000000000000001","ratio16":"1.0000000000000001"}
