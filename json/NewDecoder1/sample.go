package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

var data = `{
    "id": 1.0,
    "Name": "Google"
}`

// Data is struct for json parse.
type Data struct {
	Name string  `json:"name"`
	ID   float64 `json:"id,float64"`
}

func main() {
	fn1()
	fn2()
}

func fn1() {
	d := json.NewDecoder(strings.NewReader(data))
	d.UseNumber()
	var x interface{}
	if err := d.Decode(&x); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decoded to %#v\n", x)
	result, err := json.Marshal(x)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("encoded to %s\n", result)
}

func fn2() {
	d := json.NewDecoder(strings.NewReader(data))
	d.UseNumber()
	var x Data
	if err := d.Decode(&x); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decoded to %#v\n", x)
	result, err := json.Marshal(x)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("encoded to %s\n", result)
}

// 結果
// Data型でDecodeすると1.0が1になる
// decoded to map[string]interface {}{"Name":"Google", "id":"1.0"}
// encoded to {"Name":"Google","id":1.0}
// decoded to main.Data{Name:"Google", ID:1}
// encoded to {"name":"Google","id":1}
